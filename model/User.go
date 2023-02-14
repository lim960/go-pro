package model

import "github.com/jinzhu/gorm"

type User struct {
	//model内已经包含id，创建、修改、删除时间, 并且会在对应的操作时自动插入/更新时间
	gorm.Model
	Username string
	Password string
	Tel      string
}
