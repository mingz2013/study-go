package main

import (
	"flag"
	"fmt"
	"github.com/mingz2013/study.go/test-900-game-server/sdk"
)

var serverType *string = flag.String("t", "", "server type to boot")

func main() {
	flag.Parse()

	//flag.Usage()

	fmt.Println(*serverType)

	switch *serverType {
	case "sdk":
		sdk.Run()
	default:
		flag.Usage()
	}

}
