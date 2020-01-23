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

// ParseRows 数据 -> 数据行
func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	//  构建二维数组
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
