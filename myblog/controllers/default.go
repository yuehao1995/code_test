package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @router /portal [get]
func (c *MainController) Portal() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
