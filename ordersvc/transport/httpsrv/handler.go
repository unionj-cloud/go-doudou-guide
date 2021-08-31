package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/svc/http/model"
)

type OrdersvcHandler interface {
	PageUsers(w http.ResponseWriter, r *http.Request)
	GetHello(w http.ResponseWriter, r *http.Request)
	GetGreeting(w http.ResponseWriter, r *http.Request)
	GetHelloWorld(w http.ResponseWriter, r *http.Request)
}

func Routes(handler OrdersvcHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"PageUsers",
			"POST",
			"/page/users",
			handler.PageUsers,
		},
		{
			"Hello",
			"GET",
			"/hello",
			handler.GetHello,
		},
		{
			"Greeting",
			"GET",
			"/greeting",
			handler.GetGreeting,
		},
		{
			"HelloWorld",
			"GET",
			"/hello/world",
			handler.GetHelloWorld,
		},
	}
}
