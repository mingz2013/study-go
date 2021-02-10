package sdk

import (
	"net/http"
	"log"
	"study-go/test-900-game-server/conf"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"study-go/test-900-game-server/database"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.RemoteAddr)

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
		return
	}

	//body = string(b)
	//fmt.Println(string(b))

	var loginArgs LoginReq

	err = json.Unmarshal(b, &loginArgs)
	if err != nil {
		log.Fatal(err)
		return
	}

	//fmt.Println(loginArgs)

	// TODO db
	u := database.FindUserByDeviceId(loginArgs.DeviceId)

	c, err := conf.GetGateAddr()
	if err != nil {
		log.Fatal(err)
		return
	}

	s := c.Servers[0]

	loginRes := LoginRes{s, u}

	bodyRes, err := json.Marshal(loginRes)

	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(bodyRes)

	return
}

func Run() {

	http.HandleFunc("/login", LoginHandler)

	serverConfig, err := conf.GetSDKAddr()
	if err != nil {
		return
	}

	addr := serverConfig.Servers[0].Ip + ":" + strconv.Itoa(serverConfig.Servers[0].Port)

	//fmt.Println(addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
