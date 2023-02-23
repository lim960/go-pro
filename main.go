package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"pro/common"
	"pro/middleware"
	"pro/rabbitmq"
	"pro/task"
)

func main() {
	//初始化配置文件
	InitConfig()
	//初始化数据库
	db := common.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	//启动各模块
	Start()
	//路由配置
	r := Router()
	//端口配置
	port := viper.GetString("server.port")
	//项目启动
	panic(r.Run(":" + port))
}

func Start() {
	//初始化redis
	common.InitRedis()
	//开启定时任务
	task.Start()
	//初始化mq
	rabbitmq.InitMq()
}

func Router() *gin.Engine {
	//禁用gin输出
	gin.DefaultWriter = io.Discard
	r := gin.Default()
	//文件上传大小
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Use(
		//异常处理
		middleware.RecoveryMiddleware(),
		//跨域
		middleware.CORSMiddleware(),
		//请求参数处理
		middleware.RequestParamsMiddleware(),
		//日志
		middleware.LogMiddle(),
		//token
		middleware.TokenMiddleware(),
	)
	return StartRouter(r)
}

func InitConfig() {
	//先读取 application.yml 文件内 判断要使用哪一个环境 再读取对应配置的内容
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	//可以添加多个搜索路径  第一个找不到会找后面的
	viper.AddConfigPath("./config")
	//linux路径
	viper.AddConfigPath("/opt/config")
	//读取内容
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	mode := viper.GetString("profiles.mode")
	switch mode {
	case "dev":
		viper.SetConfigName("application-dev")
	case "tes":
		viper.SetConfigName("application-tes")
	case "pro":
		viper.SetConfigName("application-pro")
	}
	//读取内容
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	println("***********************************************")
	println("配置文件读取完成, 当前运行环境为: ", viper.GetString("name"))
	println("***********************************************")
}
