package route

import (
	"api-iris/controller"
	"github.com/kataras/iris"
	"os"
)

var Router *iris.Application

func init() {
	Router = iris.New()

	// menu
	menuController := controller.MenuController{}
	Router.Get("/menu/get/{name:string}", menuController.List)
}

func Run() {
	port := os.Getenv("port")
	if port == "" {
		port = "19494"
	}
	port = ":"+port

	Router.Run(iris.Addr(port))
}