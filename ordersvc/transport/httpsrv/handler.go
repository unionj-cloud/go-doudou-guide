package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/svc/http/model"
)

type OrdersvcHandler interface {
	PageUsers(w http.ResponseWriter, r *http.Request)
}

func Routes(handler OrdersvcHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"PageUsers",
			"POST",
			"/ordersvc/pageusers",
			handler.PageUsers,
		},
	}
}
