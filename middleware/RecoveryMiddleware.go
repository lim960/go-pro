package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				println(fmt.Sprintf("recovery异常：%s", err))
				response.Fail(c, "服务异常")
				//终止请求执行
				c.Abort()
			}
		}()
		c.Next()
	}
}
