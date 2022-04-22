package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type MenuController struct {
	BaseController
}

func (ctx *MenuController) List(c echo.Context) (err error) {
	name := c.Param("name")

	output := ResponseJson{
		Data: name,
		Error: "",
	}

	return c.JSON(http.StatusOK, output)
}
