package main

import (
	_ "study-go/test-120-beego/02-bee-tool-make/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run("localhost:8000")
}
