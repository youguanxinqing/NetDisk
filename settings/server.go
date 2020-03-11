package settings

var serverConf ServerConf

type ServerConf struct {
	Host string `ini:"host"`
	Port int32  `ini:"port"`
}

func init() {
	serverConf = _if(Global == nil, Default.Server, Global.Server).(ServerConf)
}
