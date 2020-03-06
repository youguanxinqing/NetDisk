package mysql

import (
	"database/sql"
	"log"
	"netdisk/settings"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// parseTime=true 开启解析时间
	db, _ = sql.Open(settings.DBType(), settings.DBURI())
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

// ParseRows 序列化数据
func ParseRows(rows *sql.Rows) []map[string]interface{} {
	// 获取所有列名
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	// scanArgs 指针列表; values 值列表
	for i := range values {
		scanArgs[i] = &values[i]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		CheckErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

// CheckErr 异常检查
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
