package plugins

import (
	"fmt"
)

func ExcuteWanfa(lines []map[string]interface{}) (result map[string]interface{}) {
	for i := 0; i < len(lines); i++ {
		//var item_params, play_mode string
		item_params := lines[i]["item_params"]
		play_mode := lines[i]["play_mode"]

		fmt.Println(item_params, play_mode)
		if _, ok := result[play_mode.(string)]; !ok {
			result[play_mode.(string)] = make(map[string]interface{})
		}
		for k, v := range item_params.(map[string]interface{}) {
			if k == "wamFa" {
				for i := 0; i < len(v); i++ {

				}
			} else {

			}
		}
	}
	return
}
