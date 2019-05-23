package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

// @router /query [get]
func (c *TestController) Query() {
	c.Data["json"] = "aaa"
	c.ServeJSON()
}
