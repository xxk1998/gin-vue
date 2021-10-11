package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:xxxx1111@tcp(127.0.0.1:3306)/db_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB连接失败")
	}
	return db
}

func GetDB() *gorm.DB {
	DB := InitDB()
	return DB
}
