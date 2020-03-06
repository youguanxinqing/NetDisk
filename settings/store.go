package settings

import (
	"os"
)

type StoreDirConf string

// fileConfigPath 配置文件路径
const sep = string(os.PathSeparator)
const fileConfigPath = "config.ini"

// ParseConfig 解析配置
//func ParseConfig() {
//	f, err := ini.Load(fileConfigPath)
//	if err != nil {
//		log.Println(err.Error())
//	}
//
//	sec := f.Section("Dir")
//	Global = Config{
//		// 目录尽量统一处理为: dir + '/'
//		StoreDir: parseStorePath(sec) + sep,
//	}
//}

// parseStorePath ...
//func parseStorePath(sec *ini.Section) (storeDir string) {
//	storeDir = sec.Key("STORE_DIR").String()
//
//	if finfo, err := os.Stat(storeDir); err == nil && finfo.IsDir() {
//		return
//	} else if !os.IsNotExist(err) {
//		return
//	}
//
//	os.Mkdir(storeDir, os.ModePerm)
//	return
//}
//
//func init() {
//	ParseConfig()
//}
