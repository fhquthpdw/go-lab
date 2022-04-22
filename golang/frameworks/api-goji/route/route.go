package route

import (
	"api-goji/controller"
	goji "goji.io"
	"goji.io/pat"
	"net/http"
	"os"
	"fmt"
)

var Router *goji.Mux

func init() {
	Router = goji.NewMux()

	// menu
	menuController := controller.MenuController{}
	Router.HandleFunc(pat.Get("/menu/get/:name"), menuController.List)
}

func Run() {
	port := os.Getenv("port")
	if port == "" {
		port = "19999"
	}
	port = ":"+port

    fmt.Println(port)
	http.ListenAndServe(port, Router)
}
