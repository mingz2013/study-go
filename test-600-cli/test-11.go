package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"
)

func main() {
	app := cli.NewApp()

	flags := []cli.Flag{
		altsrc.NewIntFlag(cli.IntFlag{Name: "test"}),
		cli.StringFlag{Name: "load"},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("yaml ist rad")
		return nil
	}

	app.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("load"))
	app.Flags = flags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
