package plugins

import "fmt"

func ExcuteWanfa(lines []map[string]interface{}) (result map[string]interface{}) {
	for i := 0; i < len(lines); i++ {
		var item_params, play_mode string
		item_params = lines[i]["item_params"]
		play_mode = lines[i]["play_mode"]

		fmt.Println(item_params, play_mode)
	}
}
