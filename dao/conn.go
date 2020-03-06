package dao

import (
	"netdisk/settings"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func NewDB() *gorm.DB {
	return db
}

func initDB() {
	var err error
	if db, err = gorm.Open(settings.DBType(), settings.DBURI()); err != nil {
		log.Fatal(err)
	}
}

func init() {
	initDB()
}
