package model

import (
	"fmt"
	"go_blog/utils"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db  *gorm.DB
	err error
)

func InitDb() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)

	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// 日志模式 silent
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true, //禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("connection error", err)
		os.Exit(1)
	} else {
		fmt.Println("database connection ok.")
	}
	_ = db.AutoMigrate(&User{}, &Article{}, &Category{}, &Profile{}, &Comment{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)  //连接池
	sqlDB.SetMaxOpenConns(100) //最大链接数
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
