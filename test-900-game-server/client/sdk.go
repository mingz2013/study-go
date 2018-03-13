package client

import (
	"net/http"
	"io/ioutil"
	"log"
)

func ConnectSDK(url *string) (body string, err error) {

	res, err := http.Get(*url)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	//fmt.Println(string(b))

	if err != nil {
		log.Fatal(err)
		return
	}

	body = string(b)
	//fmt.Println(body)
	return
}
