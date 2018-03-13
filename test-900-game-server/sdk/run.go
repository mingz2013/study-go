package sdk

import (
	"net/http"
	"log"
	"github.com/mingz2013/study.go/test-900-game-server/conf"
	"strconv"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
		return
	}

	//body = string(b)
	fmt.Println(string(b))

	var loginArgs LoginArgs

	err = json.Unmarshal(b, &loginArgs)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(loginArgs)
	return
}

func Run() {

	http.HandleFunc("/login", LoginHandler)

	serverConfig, err := conf.GetSDKAddr()
	if err != nil {
		return
	}

	addr := serverConfig.Servers[0].Ip + ":" + strconv.Itoa(serverConfig.Servers[0].Port)

	fmt.Println(addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
