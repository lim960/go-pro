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
		tx.AddError(tx.Create(user).Error)
		tx.AddError(tx.Create(&model.Balance{UserId: user.ID, Balance: 1000}).Error)
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
