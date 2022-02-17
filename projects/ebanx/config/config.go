package config

import (
	"os"
	"sync"
)

var conf Conf
var confOnce sync.Once

func GetConfig() Conf {
	confOnce.Do(func() {
		conf.Server = initServer()
	})
	return conf
}

type Conf struct {
	Server Server
}

type Server struct {
	Addr string
	Port string
	Mode string
}

func initServer() Server {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ""
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	model := os.Getenv("MODEL")
	if model == "" {
		model = "debug"
	}

	return Server{
		Addr: addr,
		Port: port,
		Mode: model,
	}
}
