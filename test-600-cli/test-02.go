package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"github.com/labstack/gommon/log"
)

// $ go run test-02.go
// $ go run test-02.go help

func main() {
	app := cli.NewApp()

	app.Name = "boom"

	app.Usage = "make an explosive entrance"

	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		fmt.Printf("Hello %q", c.Args().Get(0))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
