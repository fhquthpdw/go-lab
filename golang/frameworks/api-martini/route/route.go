package route

import (
	"api-martini/controller"
	"github.com/go-martini/martini"
	"os"
)

var Router *martini.ClassicMartini

func init() {
	Router = martini.Classic()

	// menu
	menuController := controller.MenuController{}
	Router.Get("/menu/get/:name", menuController.List)
}

func Run() {
	port := os.Getenv("port")
	if port == "" {
		port = "19595"
	}
	port = ":"+port

	Router.RunOnAddr(port)
}