package settings

var serverConf ServerConf

type ServerConf struct {
	Host string
	Port int32
}

func init() {
	serverConf = Default.Server
}
