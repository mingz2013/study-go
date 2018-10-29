package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/labstack/gommon/log"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:     "password, p",
			Usage:    "password for the mysql database",
			FilePath: "/etc/mysql/password",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
