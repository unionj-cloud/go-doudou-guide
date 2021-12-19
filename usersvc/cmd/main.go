package main

import (
	"fmt"
	"github.com/ascarter/requestid"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	ddconfig "github.com/unionj-cloud/go-doudou/svc/config"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
	"github.com/unionj-cloud/go-doudou/svc/registry"
	service "usersvc"
	"usersvc/config"
	"usersvc/transport/httpsrv"
)

func main() {
	ddconfig.InitEnv()
	conf := config.LoadFromEnv()

	err := registry.NewNode()
	if err != nil {
		logrus.Panicln(fmt.Sprintf("%+v", err))
	}
	//defer registry.Shutdown()

	svc := service.NewUsersvc(conf)

	handler := httpsrv.NewUsersvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddMiddleware(ddhttp.Metrics, requestid.RequestIDHandler, handlers.CompressHandler, handlers.ProxyHeaders, ddhttp.Logger, ddhttp.Rest, ddhttp.Recover)
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
