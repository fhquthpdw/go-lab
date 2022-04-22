package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MenuController struct {
	BaseController
}

func (ctx *MenuController) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	response := ResponseJson{
		Data: name,
		Error: "",
	}
	outputByte, _ := json.Marshal(response)
	outputString := string(outputByte)

	fmt.Fprint(w, outputString)
}
