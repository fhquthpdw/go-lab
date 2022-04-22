package order

import serviceContainer "ffind/service-locator/container"

type order struct {
	cacheObj serviceContainer.Cache
	msgObj   serviceContainer.Msg
}

func NewOrder(cache serviceContainer.Cache, msg serviceContainer.Msg) order {
	return order{
		cacheObj: cache,
		msgObj:   msg,
	}
}

func (o order) Send() {
	// cache
	o.cacheObj.Set("key", "value")
	o.cacheObj.Get("key")

	// msg
	o.msgObj.Send("to", "msg")
}
