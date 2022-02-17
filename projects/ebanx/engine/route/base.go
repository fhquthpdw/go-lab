package route

import (
	"ebanx/config"
	"sync"

	"github.com/gin-gonic/gin"
)

func init() {
}

var router *gin.Engine
var routerOnce sync.Once

func GetRouter() *gin.Engine {
	routerOnce.Do(func() {
		serverConf := config.GetConfig().Server
		gin.SetMode(serverConf.Mode)

		router = gin.Default()
	})

	return router
}
