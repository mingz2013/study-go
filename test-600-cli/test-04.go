package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"github.com/labstack/gommon/log"
)

// $ go run test-04.go -lang="" asd
func main() {
	var language string

	app := cli.NewApp()

	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "someone"

		if c.NArg() > 0 {
			name = c.Args()[0]
		}
		if language == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
