package engine

import (
	"ebanx/config"
	"ebanx/engine/route"
	"fmt"
	"log"
)

func Run() {
	Init()

	serverConf := config.GetConfig().Server
	servStr := fmt.Sprintf("%s:%s", serverConf.Addr, serverConf.Port)
	if err := route.GetRouter().Run(servStr); err != nil {
		log.Fatalf("application run error: %s", err)
	}
}

func Init() {
	for _, f := range []func(){
		InitRouters,
	} {
		f()
	}
}

func InitRouters() {
	route.AccountRoutes()
	route.SysRoutes()
}
