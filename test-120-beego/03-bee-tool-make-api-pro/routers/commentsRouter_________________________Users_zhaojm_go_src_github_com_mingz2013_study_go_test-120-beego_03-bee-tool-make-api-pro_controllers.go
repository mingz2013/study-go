package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingz2013/study.go/test-120-beego/03-bee-tool-make-api-pro/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
