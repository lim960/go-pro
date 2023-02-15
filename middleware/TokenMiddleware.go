package middleware

import (
	"github.com/gin-gonic/gin"
	"pro/response"
	"strings"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//检查放行白名单
		url := c.Request.RequestURI
		if CheckWhitelist(url) {
			return
		}
		//验证token
		token := c.GetHeader("token")
		if CheckToken(token) {
			return
		}
		response.Fail(c, "token失效")
		//终止请求执行
		c.Abort()
	}
}

// 放行白名单 配置具体接口/群组
var Whitelist = []string{
	"/user/ggt",
	"/rel/",
}

// 检查接口是否需要放行
func CheckWhitelist(url string) bool {

	for _, item := range Whitelist {
		if url == item || (strings.HasSuffix(item, "/") && strings.HasPrefix(url, item)) {
			return true
		}
	}
	return false
}

// 校验token  token校验及权限校验在此处处理
func CheckToken(token string) bool {

	if token == "123" {
		return true
	}
	return false
}
