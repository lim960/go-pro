package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"pro/response"
	"strings"
)

var prefix string

// Whitelist 放行白名单
var Whitelist []string

func InitWhite() {
	Whitelist = []string{
		prefix + "/common/",
	}
}

func TokenMiddleware() gin.HandlerFunc {
	prefix = viper.GetString("server.prefix")
	InitWhite()
	return func(c *gin.Context) {
		//检查放行白名单
		url := c.Request.RequestURI
		if CheckWhitelist(url) {
			return
		}
		//验证token
		token := c.GetHeader("token")
		if token == "" {
			Fail(c, "token无效")
			return
		}
		if strings.HasPrefix(url, prefix+"/back") {
			//后台接口
			pass, msg := CheckBackToken(c, token, url)
			if !pass {
				Fail(c, msg)
				return
			}
		} else if strings.HasPrefix(url, prefix+"/app") {
			//app接口
			pass, msg := CheckAppToken(c, token, url)
			if !pass {
				Fail(c, msg)
				return
			}
		}
	}
}

// CheckBackToken 后台接口鉴权
func CheckBackToken(c *gin.Context, token, url string) (bool, string) {

	if token == "123" {
		return true, "ok"
	}
	return false, "权限不足"
}

// CheckAppToken app接口鉴权
func CheckAppToken(c *gin.Context, token, url string) (bool, string) {

	if token == "123" {
		return true, "ok"
	}
	return false, "权限不足"
}

// CheckWhitelist 检查接口是否需要放行
func CheckWhitelist(url string) bool {
	for _, item := range Whitelist {
		if url == item || (strings.HasSuffix(item, "/") && strings.HasPrefix(url, item)) {
			return true
		}
	}
	return false
}

func Fail(c *gin.Context, msg string) {
	response.TokenFail(c, msg)
	//终止请求执行
	c.Abort()
}
