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
			Name:   "lang, l",
			Value:  "english",
			Usage:  "language for the greeting",
			EnvVar: "LEGACY_COMPAT_LANG,APP_LANG,LANG",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
