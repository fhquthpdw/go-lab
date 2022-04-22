package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MenuController struct {
	BaseController
}

func (ctx *MenuController) List(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	response := ResponseJson{
		Data: name,
		Error: "",
	}
	outputByte, _ := json.Marshal(response)
	outputString := string(outputByte)

	fmt.Fprint(w, outputString)
}
