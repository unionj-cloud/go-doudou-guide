package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	service "statsvc"
	"statsvc/config"
	"statsvc/internal/reportsvcj"
	"statsvc/transport/httpsrv"
)

func main() {
	conf := config.LoadFromEnv()

	err := registry.NewNode()
	if err != nil {
		logrus.Panic(fmt.Sprintf("%+v", err))
	}
	defer registry.Shutdown()

	svc := service.NewStatsvc(conf,
		reportsvcj.NewEcho(
			ddhttp.WithRootPath("/report-svc-j"),
			ddhttp.WithProvider(ddhttp.NewNacosWRRServiceProvider("report-svc-j"))),
	)
	handler := httpsrv.NewStatsvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
