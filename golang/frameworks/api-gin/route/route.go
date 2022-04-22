package route

import (
	"api-gin/controller"
	"github.com/gin-gonic/gin"
	"os"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	// menu
	menuController := controller.MenuController{}
	Router.GET("/menu/get/:name", menuController.List)
}

func Run() {
	port := os.Getenv("port")
	if port == "" {
		port = "19393"
	}
	port = ":"+port

	Router.Run(port)
}
