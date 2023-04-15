package controller

import (
	"github.com/gin-gonic/gin"
	"pro/model"
	"pro/response"
	"pro/service"
	"strconv"
)

func SaveUser(c *gin.Context) {
	//获取参数 a b
	username := c.GetString("username")
	id := uint(c.GetFloat64("id"))
	//验证此处有没有接收到参数
	//println("username", username)
	//println("id", id)
	user := model.User{
		Username: username,
	}
	service.Saveuser(id, &user)
	response.Success(c, user)
}

func DelUser(c *gin.Context) {
	//获取参数 a b
	id, _ := strconv.Atoi(c.Query("id"))

	service.DelUser(uint(id))
	response.Success(c, nil)
}

func GetUser(c *gin.Context) {

	//获取参数 a b
	id, _ := strconv.Atoi(c.Query("id"))

	var user model.User
	service.GetUser(uint(id), &user)
	response.Success(c, user)
}
