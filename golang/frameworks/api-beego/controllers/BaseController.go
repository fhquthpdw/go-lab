package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

type ResponseJson struct {
	Data  interface{} `json:"data"`
	Error string `json:"error"`
}
