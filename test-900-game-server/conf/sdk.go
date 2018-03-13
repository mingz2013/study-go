package conf

import "fmt"

type ServerJson struct {
	Ip   string "ip"
	Port int    "port"
}

func GetSDKAddr() (ip string, port int, err error) {

	var v ServerJson

	err = LoadJson("config/sdk/0.json", &v)
	if err != nil {
		return
	}
	fmt.Println(v)
	ip = v.Ip
	port = v.Port
	return
}
