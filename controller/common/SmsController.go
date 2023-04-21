package common

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"pro/common"
	"pro/consts"
	"pro/response"
	"pro/util"
	"pro/util/third"
)

// SendSms 发送短信验证码
func SendSms(ctx *gin.Context) {

	tel := ctx.GetString("tel")
	// 0-用户注册 1-用户登录 2-修改密码 3- 4- 5-商户注册 6-商户登录
	types := int(ctx.GetFloat64("types"))

	if tel == "" {
		panic("参数错误")
	}
	env := viper.GetString("server.env")
	key := consts.SmsKey[types] + tel
	code := "1234"
	if env == "生产环境" {
		code = util.GetCode()
		third.Send(tel, code)
	}
	println(tel + " 验证码为: " + code)
	//缓存
	common.SetWithTime(key, code, 5*60)
	response.Success(ctx, "success")
}
