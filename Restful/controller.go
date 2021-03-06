package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db            *gorm.DB
	sqlConnection = "root:shutian1.2@localhost/demo?charset=utf8&parseTime=True&loc=Local"
)

func init() {
	var err error
	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&todoModel{})
}
