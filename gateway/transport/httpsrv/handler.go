package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type GatewayHandler interface {
	PageUsers(w http.ResponseWriter, r *http.Request)
}

func Routes(handler GatewayHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"PageUsers",
			"POST",
			"/page/users",
			handler.PageUsers,
		},
	}
}
