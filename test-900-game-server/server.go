package main

import (
	"github.com/mingz2013/study.go/test-900-game-server/sdk"
	"github.com/mingz2013/study.go/test-900-game-server/gate"
)

//var serverType *string = flag.String("t", "", "server type to boot")

func main() {
	//flag.Parse()

	//flag.Usage()

	//fmt.Println(*serverType)
	//
	//switch *serverType {
	//case "sdk":
	//	sdk.Run()
	//default:
	//	flag.Usage()
	//}
	ch := make(chan string, 0)

	//go agent.Run()

	//go game.Run()

	go gate.Run()

	go sdk.Run()

	//go http.Run()

	<-ch

}
