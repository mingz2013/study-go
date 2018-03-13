package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/mingz2013/study.go/test-900-game-server/client"
	"github.com/mingz2013/study.go/test-900-game-server/conf"
	"encoding/json"
)

var url *string = flag.String("url", "http://localhost:8000/login", "sdk url")

func main() {
	flag.Parse()

	body, err := client.ConnectSDK(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(body)

	var gateAddr conf.ServerAddr
	err = json.Unmarshal(body, &gateAddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	conn, err := client.ConnectGate(gateAddr.Ip, gateAddr.Port)

	if err != nil {
		log.Fatal(err)
		return
	}

	client.HandleConn(conn)
}
