package main

import (
	"fmt"
	service "gateway"
	"gateway/config"
	"gateway/transport/httpsrv"
	"github.com/sirupsen/logrus"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
)

func main() {
	err := registry.NewNode()
	if err != nil {
		logrus.Panic(fmt.Sprintf("%+v", err))
	}
	defer registry.Shutdown()

	conf := config.LoadFromEnv()
	svc := service.NewGateway(conf)
	handler := httpsrv.NewGatewayHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddMiddleware(ddhttp.Proxy(ddhttp.ProxyConfig{}))
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
