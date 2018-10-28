package routers

import (
	"github.com/mingz2013/study.go/test-120-beego/02-bee-tool-make/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
