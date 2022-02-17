package route

import "ebanx/api"

const (
	Event   = "/event"
	Balance = "/balance"
)

// AccountRoutes AccountRouters contains all the routes for the account
func AccountRoutes() {
	r := GetRouter()
	ctl := api.Account{}

	r.GET(Balance, ctl.Balance)
	r.POST(Event, ctl.Event)
}
