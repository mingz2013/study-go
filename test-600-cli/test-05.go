package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/labstack/gommon/log"
)

// GLOBAL OPTIONS:
//    --config FILE, -c FILE  Load configuration from FILE

//

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
