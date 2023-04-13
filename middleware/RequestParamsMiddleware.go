package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
)

//请求参数处理

func RequestParamsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(
			strings.ToLower(c.Request.Header.Get("Content-Type")),
			"application/json",
		) {
			//读取全部入参
			var params map[string]any
			json.NewDecoder(c.Request.Body).Decode(&params)
			//遍历入参 set存储 便于单一方式取参
			for key := range params {
				c.Set(key, params[key])
			}
			b, _ := json.Marshal(params)
			//存取完整参数 便于打印请求参数
			c.Set("reqParam", params)
			println(fmt.Sprintf("reqParam: %s", string(b)))
			//参数重新放回body  便于绑定取参
			c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		}

	}
}
