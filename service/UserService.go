package service

import (
	"pro/common"
	"pro/model"
)

func Saveuser(id uint, user *model.User) {

	db := common.GetDB()
	if id > 0 {
		db.Model(&model.User{}).Where("id = ?", id).Updates(user)
	} else {
		//开启事务
		tx := db.Begin()
		defer AutoTransaction(tx)
		tx = tx.Create(user)
		tx = tx.Create(&model.Balance{UserId: user.ID, Balanc: 1000})
	}
}

func DelUser(id uint) {

	db := common.GetDB()

	db.Delete(&model.User{}, id)
}

func GetUser(id uint, user *model.User) {

	db := common.GetDB()

	db.Where("id = ?", id).Find(user)
}
