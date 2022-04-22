package routers

import (
	"api-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/menu/get/:name", &controllers.MenuController{}, "get:List")
}
