package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func init() {
	// parseTime=true 开启解析时间
	db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/netdisk?charset=utf8&parseTime=true")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		log.Println("Failed connect mysql, err:" + err.Error())
		os.Exit(1)
	}
}

// DBConn 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
