package main

import (
	"github.com/mingz2013/study.go/test-900-game-server/sdk"
	"github.com/mingz2013/study.go/test-900-game-server/gate"
	"github.com/mingz2013/study.go/test-900-game-server/agent"
	"github.com/mingz2013/study.go/test-900-game-server/game"
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

	go agent.Run()

	go game.Run()

	go gate.Run()

	go sdk.Run()

	<-ch

}
