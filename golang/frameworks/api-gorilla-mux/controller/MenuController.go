package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MenuController struct {
	BaseController
}

func (ctx *MenuController) List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	response := ResponseJson{
		Data: name,
		Error: "",
	}
	outputByte, _ := json.Marshal(response)
	outputString := string(outputByte)

	fmt.Fprint(w, outputString)
}
