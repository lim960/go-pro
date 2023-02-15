package main

import (
	"github.com/gin-gonic/gin"
	"pro/controller"
)

func StartRouter(r *gin.Engine) *gin.Engine {
	//用户
	user := r.Group("/user")
	{
		user.POST("/save", controller.SaveUser)
		user.POST("/del", controller.DelUser)
		user.POST("/get", controller.GetUser)
		user.POST("/ggt", controller.GetUser)
		user.GET("/gg", controller.GG)
	}

	rel := r.Group("/rel")
	{
		rel.POST("/get", controller.GetUser)
	}

	return r
}
