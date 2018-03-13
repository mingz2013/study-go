package sdk

import (
	"net/http"
	"log"
	"github.com/mingz2013/study.go/test-900-game-server/conf"
	"strconv"
	"fmt"
)

func Run() {

	var s SDK

	serverConfig, err := conf.GetSDKAddr()
	if err != nil {
		return
	}

	addr := serverConfig.Servers[0].Ip + ":" + strconv.Itoa(serverConfig.Servers[0].Port)

	fmt.Println(addr)

	err = http.ListenAndServe(addr, s)
	if err != nil {
		log.Fatal(err)
	}
}
