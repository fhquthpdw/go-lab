package main

import (
	serviceContainer "ffind/service-locator/container"
	"ffind/service-locator/order"
)

func main() {
	cacheIns := serviceContainer.Services.GetService(serviceContainer.CacheSvcToken).(serviceContainer.Cache)
	msgIns := serviceContainer.Services.GetService(serviceContainer.MsgSvcToken).(serviceContainer.Msg)
	o := order.NewOrder(cacheIns, msgIns)
	o.Send()
}
