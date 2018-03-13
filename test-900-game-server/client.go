package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/mingz2013/study.go/test-900-game-server/client"
)

var url *string = flag.String("url", "http://localhost:8000/login", "sdk url")

func main() {
	flag.Parse()

	body, err := client.ConnectSDK(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body)

}
