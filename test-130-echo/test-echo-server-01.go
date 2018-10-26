package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "hello, world!")
	})
	e.Logger.Fatal(e.Start(":8001"))
}
