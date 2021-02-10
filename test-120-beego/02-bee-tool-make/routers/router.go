package routers

import (
	"study-go/test-120-beego/02-bee-tool-make/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
