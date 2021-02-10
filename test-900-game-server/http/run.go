package http

import (
	"net/http"
	"strconv"
	"fmt"
	"log"
	"study-go/test-900-game-server/conf"
)

func ManagerHandler(w http.ResponseWriter, r *http.Request) {

}

func Run() {

	http.HandleFunc("/manager", ManagerHandler)

	serverConfig, err := conf.GetHTTPAddr()
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
