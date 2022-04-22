package controller

type BaseController struct {
	//
}

type ResponseJson struct {
	Data  interface{} `json:"data"`
	Error string `json:"error"`
}
