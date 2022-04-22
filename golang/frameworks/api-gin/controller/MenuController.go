package controller

import (
	"github.com/gin-gonic/gin"
)

type MenuController struct {
	BaseController
}

func (ctx *MenuController) List(c *gin.Context) {
	name := c.Param("name")

	response := ResponseJson{
		Data: name,
		Error: "",
	}

	c.JSON(200, response)
}
