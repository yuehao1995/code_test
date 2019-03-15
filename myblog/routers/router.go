// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"github.com/astaxie/beego"
	"github.com/eechains/code_test/myblog/controllers"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/test",
				beego.NSInclude(
					&controllers.MainController{},
				),
			),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
		)
	beego.AddNamespace(ns)
	beego.InsertFilter("/*", beego.BeforeRouter, ValidateMiddlerWare)
}
