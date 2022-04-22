package controller

import (
	"encoding/json"
	"github.com/go-martini/martini"
)

type MenuController struct {
	BaseController
}

func (ctx *MenuController) List(p martini.Params) string {
	name := p["name"]

	output := ResponseJson{
		Data: name,
		Error: "",
	}
	outputByte, _ := json.Marshal(output)
	outputStr := string(outputByte)

	return outputStr
}
