package log_analysis

import (
	"fmt"
	"github.com/mingz2013/study.go/log_analysis/utils"
	"encoding/json"
	"github.com/mingz2013/study.go/log_analysis/plugins"
)

func main() {
	day_list := []string{"2017_07_04", "2017_07_05"}
	for i := 0; i < len(day_list); i++ {
		do_day(day_list[i])
	}

}

func do_day(day string) {

	file_name := "data/analysis_" + day + ".json"

	fmt.Println(file_name)

	var l []map[string]interface{}
	l = get_list(file_name)
	fmt.Println(l)
	var result map[string]interface{}

	result = plugins.ExcuteWanfa(l)
	fmt.Println(result)
	utils.WriteObjToJsonFile(result, "result/wanfa_"+day+".json")



}

func get_list(file_name string) (lines2 []map[string]interface{}) {
	lines := utils.ReadFileLineToList(file_name)
	for i := 0; i < len(lines); i++ {
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(lines[0]), &m); err == nil {
			lines2 = append(lines2, m)
		}

	}
	return
}
