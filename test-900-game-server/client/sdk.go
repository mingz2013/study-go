package client

import (
	"net/http"
	"io/ioutil"
	"log"
	"github.com/mingz2013/study.go/test-900-game-server/sdk"
	"encoding/json"
	"bytes"
)

func ConnectSDK(url *string) (ret []byte, err error) {

	loginBody := &sdk.LoginArgs{DeviceId: DEVICE_ID}

	loginBodyJson, err := json.Marshal(loginBody)

	if err != nil {
		log.Fatal(err)
		return
	}
	body := bytes.NewReader(loginBodyJson)

	res, err := http.Post(*url, "json", body)

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

	ret = b
	//fmt.Println(body)
	return
}
