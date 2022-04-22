package route

import (
	"api-gorilla-mux/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	// menu
	menuController := controller.MenuController{}
	Router.HandleFunc("/menu/get/{name}", menuController.List).Methods("GET")
}

func Run() {
	port := os.Getenv("port")
	if port == "" {
		port = "19797"
	}
	port = ":"+port

	http.ListenAndServe(port, Router)
}
