package httpsrv

import (
	"net/http"

	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
)

type AdminHandler interface {
	PageUsers(w http.ResponseWriter, r *http.Request)
}

func Routes(handler AdminHandler) []ddhttp.Route {
	return []ddhttp.Route{
		{
			"PageUsers",
			"POST",
			"/admin/pageusers",
			handler.PageUsers,
		},
	}
}
