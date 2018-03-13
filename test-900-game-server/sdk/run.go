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

	ip, port, err := conf.GetSDKAddr()
	if err != nil {
		return
	}

	addr := ip + ":" + strconv.Itoa(port)

	fmt.Println(addr)

	err = http.ListenAndServe(addr, s)
	if err != nil {
		log.Fatal(err)
	}
}
