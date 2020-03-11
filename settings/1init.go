package settings

import (
	"os"

	log "github.com/sirupsen/logrus"

	"gopkg.in/ini.v1"
)

type Config struct {
	Store    StoreConf
	DataBase DataBaseConf
	Server   ServerConf
}

var (
	// 项目使用的 Config 实例
	Global *Config
	// 存储默认 Config 的实例，当 config.ini 不存在时，将其赋值
	// 给 Global，得以让项目使用
	Default *Config
)

const (
	configFilePath = "config.ini"
)

func prepare() {
	// 1. 加载配置文件
	f, err := ini.Load(configFilePath)
	if err != nil {
		log.Warn("this is no config.ini, but don't worry if you run first")
	} else {
		// 1.1 加载成功后映射
		c := new(Config)
		if err := f.MapTo(c); err != nil {
			log.Error(err)
			os.Exit(0)
		}
		Global = c
		return
	}

	// 2. 若加载失败则生成默认配置
	cfg := ini.Empty()
	err = ini.ReflectFrom(cfg, Default)
	if err != nil {
		log.Error(err)
		os.Exit(0)
	}

	// 3. 保存至本地
	err = cfg.SaveTo("config.ini")
	if err != nil {
		log.Error(err)
		os.Exit(0)
	}
}

func init() {
	Default = &Config{
		Store: StoreConf{
			StoreDir: "/tmp/netdisk",
		},
		DataBase: DataBaseConf{
			Type:     "mysql",
			Host:     "guan.com",
			Port:     3306,
			UserName: "root",
			Password: "123456",
			DBName:   "netdisk",
		},
		Server: ServerConf{
			Host: "0.0.0.0",
			Port: 8080,
		},
	}

	// 预处理
	prepare()

	if Global == nil {
		Global = Default
	}
}
