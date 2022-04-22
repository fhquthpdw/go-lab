package route

import (
	"api-echo/controller"
	"github.com/labstack/echo/v4"
	"os"
)

var Router *echo.Echo

func init() {
	Router = echo.New()

	// menu
	menuController := controller.MenuController{}
	Router.GET("/menu/get/:name", menuController.List)
}

func Run() {
	port := os.Getenv("port")
	if port == "" {
		port = "19191"
	}
	port = ":"+port

	Router.Logger.Fatal(Router.Start(port))
}
