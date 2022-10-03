package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	DSN := "root:password@(localhost)/simplerest?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", DSN)
	if err == nil {
		db = d
	} else {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
