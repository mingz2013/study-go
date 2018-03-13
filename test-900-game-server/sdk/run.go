package sdk

import (
	"net/http"
	"log"
)

func Run() {

	var s SDK

	err := http.ListenAndServe("localhost:8000", s)
	if err != nil {
		log.Fatal(err)
	}
}
