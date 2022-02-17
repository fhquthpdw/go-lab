package route

import "ebanx/api"

const (
	Reset = "/reset"
)

// SysRoutes is a collection of routes for the sys endpoints
func SysRoutes() {
	r := GetRouter()
	ctl := api.Sys{}

	r.POST(Reset, ctl.Reset)
}
