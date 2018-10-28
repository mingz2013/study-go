package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello beego")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
