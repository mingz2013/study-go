package conf

type ServerAddr struct {
	Ip   string
	Port int
}

type ServerConfig struct {
	Servers []ServerAddr
}
