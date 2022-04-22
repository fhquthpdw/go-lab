package route

import (
	"api-httprouter/controller"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"fmt"
)

var Router *httprouter.Router

func init() {
	Router = httprouter.New()

	// menu
	menuController := controller.MenuController{}
	Router.GET("/menu/get/:name", menuController.List)
}

func Run() {
	port := os.Getenv("port")
	if port == "" {
		port = "19292"
	}
	port = ":"+port

    fmt.Println(port)
	log.Fatal(http.ListenAndServe(port, Router))
}
