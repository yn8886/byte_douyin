package models

import (


	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	dsn := "root:020310yn@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local" 
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{}, &UserLogin{})
	if err != nil {
		panic(err)
	}
	DB = db
}