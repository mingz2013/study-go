package conf

import "fmt"

func GetGateAddr() (ServerConfig, error) {

	var c ServerConfig

	err := LoadJsonFile("config/gate/0.json", &c)
	if err != nil {
		return c, err
	}
	fmt.Println(c)
	return c, nil
}
