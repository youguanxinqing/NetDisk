package settings

import (
	"fmt"
	"log"
)

func DBURI() string {
	t := Global.DataBase.Type

	var uri string
	switch t {
	case "mysql":
		uri = mysqlUri()
	default:
		log.Fatal("暂不支持" + t + "类型的数据库")
	}
	return uri
}

func DBType() string {
	return databaseConf.Type
}

func ServerAddr() string {
	return fmt.Sprintf(":%d", serverConf.Port)
}
