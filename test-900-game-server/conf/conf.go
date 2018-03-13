package conf

import "fmt"

type ServerAddr struct {
	Ip   string
	Port int
}

type ServerConfig struct {
	Servers []ServerAddr
}

func GetServerAddr(filename string) (ServerConfig, error) {

	var c ServerConfig

	err := LoadJsonFile(filename, &c)
	if err != nil {
		return c, err
	}
	fmt.Println(c)
	return c, nil
}

func GetSDKAddr() (ServerConfig, error) {
	return GetServerAddr("config/sdk/0.json")
}

func GetGateAddr() (ServerConfig, error) {
	return GetServerAddr("config/gate/0.json")
}

func GetAgentAddr() (ServerConfig, error) {
	return GetServerAddr("config/agent/0.json")
}

func GetGameAddr() (ServerConfig, error) {
	return GetServerAddr("config/game/0.json")
}

func GetHTTPAddr() (ServerConfig, error) {
	return GetServerAddr("config/http/0.json")
}

