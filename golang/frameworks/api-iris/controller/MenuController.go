package controller

import (
	"github.com/kataras/iris"
)

type MenuController struct {
	BaseController
}

func (ctx *MenuController) List(c iris.Context) {
	name := c.Params().Get("name")

	output := ResponseJson{
		Data: name,
		Error: "",
	}

	c.JSON(output)
}
