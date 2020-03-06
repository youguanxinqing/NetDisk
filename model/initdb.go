package model

import (
	"netdisk/dao"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func createTable(table interface{}) {
	if !db.HasTable(table) {
		db.CreateTable(table)
	} else {
		db.AutoMigrate(table)
	}
}

func init() {
	db = dao.NewDB()

	createTable(&UserModel{})
}
