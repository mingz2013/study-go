package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/labstack/gommon/log"
)

// $ go run test-01.go

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
