package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type ServerJson struct {
	Ip   string
	Port int
}

func main() {

	s := []byte(`{"ip": "localhost", "port": 8000}`)

	var serverJson ServerJson

	err := json.Unmarshal(s, &serverJson)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(serverJson)
}
