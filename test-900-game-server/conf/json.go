package conf

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func LoadJsonFile(filename string, v interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	//datajson := []byte(data)

	fmt.Println(v)

	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}

	fmt.Println(v)
	return
}
