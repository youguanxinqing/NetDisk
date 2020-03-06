package settings

type Config struct {
	StoreDir StoreDirConf
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

func init() {
	Default = &Config{
		StoreDir: "/tmp/netdisk",
		DataBase: DataBaseConf{
			Type:     "mysql",
			Host:     "192.168.111.136",
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

	if Global == nil {
		Global = Default
	}
}
