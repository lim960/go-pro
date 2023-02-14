package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mattn/go-colorable"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

// 实例化
var InfoLog *logrus.Logger
var ErrLog *logrus.Logger

func init() {
	infoName := "info.log"
	InfoLog = initLog(infoName)
	errName := "err.log"
	ErrLog = initLog(errName)
}

func initLog(fileName string) *logrus.Logger {
	fileName = "./logs/" + fileName
	// 写入文件
	var f *os.File
	var err error
	//判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		f, err = os.Create(fileName)
	} else {
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
	if err != nil {
		fmt.Println("open log file failed")
	}
	//初始化
	log := logrus.New()
	//设置日志级别
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true, //色彩启用
		FullTimestamp: true, //时间戳
		//DisableLevelTruncation: true, //禁用截断
		PadLevelText: true, //宽度相同
	})
	//同时输出到控制台和文件
	writers := []io.Writer{
		f,
		colorable.NewColorableStdout(),
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	//设置输出
	log.SetOutput(fileAndStdoutWriter)

	//log.SetOutput(f)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	log.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))
	//log.AddHook(lfshook.NewHook(writeMap, &logrus.TextFormatter{
	//	ForceColors:            true,
	//	FullTimestamp:          true,
	//	DisableLevelTruncation: true,
	//}))
	return log
}

// 日志
func LogMiddle() gin.HandlerFunc {

	return func(c *gin.Context) {

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		//statusCode := c.Writer.Status()
		//请求ip
		//clientIP := c.ClientIP()
		//请求参数
		reqParam, _ := c.Get("reqParam")
		//响应数据
		responseBody := bodyLogWriter.body.String()
		var responseData interface{}
		if responseBody != "" {
			res := Result{}
			err := json.Unmarshal([]byte(responseBody), &res)
			if err == nil {
				responseData = res.Data
			}
		}
		// 日志格式
		InfoLog.WithFields(logrus.Fields{
			//"status_code":  statusCode,
			//"client_ip":    clientIP,
			"req_uri":      reqUrl,
			"req_method":   reqMethod,
			"req_param":    reqParam,
			"resp":         responseData,
			"latency_time": latencyTime,
		}).Info()
	}
}

func Err(obj ...any) {
	ErrLog.Error(obj)
}

func Info(obj ...any) {
	InfoLog.Info(obj)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

type Result struct {
	Code  int    `json:"code"`
	Data  any    `json:"data"`
	Error string `json:"error"`
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
