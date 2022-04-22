package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"goji.io/pat"
)

type MenuController struct {
	BaseController
}

func (ctx *MenuController) List(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")

	response := ResponseJson{
		Data: name,
		Error: "",
	}
	outputByte, _ := json.Marshal(response)
	outputString := string(outputByte)

	fmt.Fprint(w, outputString)
}
