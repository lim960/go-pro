package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"pro/controller/common"
)

var prefix = viper.GetString("prefix")

func StartRouter(r *gin.Engine) *gin.Engine {
	r = CommonRouter(r)
	r = AppRouter(r)
	r = BackRouter(r)
	return r
}

// CommonRouter 公用开放接口
func CommonRouter(r *gin.Engine) *gin.Engine {
	commonPath := prefix + "/common"
	//文件上传
	file := r.Group(commonPath + "/file")
	{
		file.POST("/upload", common.Upload)
		file.POST("/batchUpload", common.BatchUpload)
	}
	//短信
	sms := r.Group(commonPath + "/sms")
	{
		sms.POST("/send", common.SendSms)
	}
	return r
}

// AppRouter app接口
func AppRouter(r *gin.Engine) *gin.Engine {
	//appPath := prefix + "/app"
	//文件上传
	//file := r.Group(appPath + "/file")
	//{
	//	file.POST("/upload", shop.Upload)
	//	file.POST("/batchUpload", shop.BatchUpload)
	//}

	return r
}

// BackRouter 后台接口
func BackRouter(r *gin.Engine) *gin.Engine {
	//backPath := prefix + "/back"
	//文件上传
	//file := r.Group(backPath + "/file")
	//{
	//	file.POST("/upload", shop.Upload)
	//	file.POST("/batchUpload", shop.BatchUpload)
	//}

	return r
}
