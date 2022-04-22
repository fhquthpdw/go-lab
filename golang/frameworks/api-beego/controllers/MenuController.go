package controllers

import (
	"github.com/astaxie/beego"
)

type MenuController struct {
	beego.Controller
}

func (c *MenuController) List() {
	name := c.Ctx.Input.Param(":name")

	output := ResponseJson{
		Data: name,
		Error: "",
	}

	c.Data["json"] = output

	c.ServeJSON()
}
