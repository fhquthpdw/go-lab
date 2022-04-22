package route

import (
	"api-net-http/controller"
	"log"
	Router "net/http"
	"os"
)

func init() {
	menuController := controller.MenuController{}
	Router.Handle("/menu/get", Router.HandlerFunc(menuController.List))
}

func Run() {
	port := os.Getenv("port")
	if port == "" {
		port = "19898"
	}
	port = ":"+port

	log.Fatal(Router.ListenAndServe(port, nil))
}
