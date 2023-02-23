package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"net/url"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 跳过默认事务
		SkipDefaultTransaction: true,
		//表设置
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // 表名前缀
			SingularTable: true,  // 禁用复数表
			NoLowerCase:   false, // 启用驼峰命名
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,       // 禁用彩色打印
			},
		),
	})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	sqlDB, err := db.DB()
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 最大连接时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
