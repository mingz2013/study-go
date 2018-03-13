package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/mingz2013/study.go/test-900-game-server/client"
	"encoding/json"
	"github.com/mingz2013/study.go/test-900-game-server/sdk"
	"github.com/mingz2013/study.go/test-900-game-server/database"
)

var url *string = flag.String("url", "http://localhost:8000/login", "sdk url")

var MyUser database.User

func main() {
	flag.Parse()

	body, err := client.ConnectSDK(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(body)

	var loginRes sdk.LoginRes
	err = json.Unmarshal(body, &loginRes)
	if err != nil {
		log.Fatal(err)
		return
	}

	MyUser = loginRes.User

	conn, err := client.ConnectGate(loginRes.ServerAddr.Ip, loginRes.ServerAddr.Port)

	if err != nil {
		log.Fatal(err)
		return
	}

	client.HandleConn(conn)
}
