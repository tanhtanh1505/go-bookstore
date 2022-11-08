package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "root:password@/go-bookstore?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bookstore")
	if err != nil {
		panic("failed to connect database")
	}
	print("connected to database")
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
