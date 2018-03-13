package conf

import "fmt"

type ServerJson struct {
	ip   string "ip"
	port int    "port"
}

func GetSDKAddr() (ip string, port int, err error) {

	v := ServerJson{}

	err = LoadJson("config/sdk/0.json", &v)
	if err != nil {
		return
	}
	fmt.Println(v)
	ip = v.ip
	port = v.port
	return
}
