package httpsrv

import (
	"net/http"

	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
)

type OrdersvcHandler interface {
	PageUsers(w http.ResponseWriter, r *http.Request)
}

func Routes(handler OrdersvcHandler) []ddhttp.Route {
	return []ddhttp.Route{
		{
			"PageUsers",
			"POST",
			"/ordersvc/pageusers",
			handler.PageUsers,
		},
	}
}
