package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data any, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data any) {
	Response(ctx, http.StatusOK, 200, data, "success")
}

func Fail(ctx *gin.Context, msg string) {
	Response(ctx, http.StatusOK, 400, nil, msg)
}

func TokenFail(ctx *gin.Context, msg string) {
	Response(ctx, http.StatusOK, -1, nil, msg)
}
