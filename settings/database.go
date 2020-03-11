package settings

import "fmt"

type DataBaseConf struct {
	Type     string `ini:"type"` // 数据库类型
	Host     string `ini:"host"`
	Port     uint16 `ini:"port"`
	UserName string `ini:"userName"`
	Password string `ini:"password"`
	DBName   string `ini:"dbname"` // 数据库名
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
	databaseConf = _if(Global == nil, Default.DataBase, Global.DataBase).(DataBaseConf)
}
