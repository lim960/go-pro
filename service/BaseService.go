package service

import (
	"gorm.io/gorm"
	"pro/middleware/log"
)

// AutoTransaction 自动处理事务
func AutoTransaction(tx *gorm.DB) {
	//如果存在数据库操作异常 回滚事务
	if tx.Error != nil {
		log.Err(tx.Error.Error())
		tx.Rollback()
		return
	}
	//如果出现运行时异常  回滚事务
	//若无异常 提交事务
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	} else {
		tx.Commit()
	}
}
