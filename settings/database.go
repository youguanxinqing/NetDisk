package settings

import "fmt"

type DataBaseConf struct {
	Type     string // 数据库类型
	Host     string
	Port     uint16
	UserName string
	Password string
	DBName   string // 数据库名
}

var databaseConf DataBaseConf

func mysqlUri() string {
	// template: "root:123456@tcp(127.0.0.1:3306)/netdisk?charset=utf8&parseTime=true"
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		databaseConf.UserName, databaseConf.Password,
		databaseConf.Host, databaseConf.Port,
		databaseConf.DBName,
	)
}

func init() {
	databaseConf = Default.DataBase
}
