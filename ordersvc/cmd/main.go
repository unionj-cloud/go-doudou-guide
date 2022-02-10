package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	"github.com/unionj-cloud/go-doudou/framework/tracing"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	service "ordersvc"
	"ordersvc/config"
	"ordersvc/transport/httpsrv"
	"os"
	"path/filepath"
	"usersvc/client"
)

func main() {
	conf := config.LoadFromEnv()

	if logger.CheckDev() {
		logger.Init(logger.WithWritter(os.Stdout))
	} else {
		logger.Init(logger.WithWritter(io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename:   filepath.Join(os.Getenv("LOG_PATH"), fmt.Sprintf("%s.log", "ordersvc")),
			MaxSize:    5,  // Max megabytes before log is rotated
			MaxBackups: 10, // Max number of old log files to keep
			MaxAge:     7,  // Max number of days to retain log files
			Compress:   true,
		})))
	}

	err := registry.NewNode()
	if err != nil {
		logrus.Panicln(fmt.Sprintf("%+v", err))
	}
	defer registry.Shutdown()

	tracer, closer := tracing.Init()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	usersvcProvider := ddhttp.NewSmoothWeightedRoundRobinProvider("usersvc")
	usersvcClient := client.NewUsersvcClient(ddhttp.WithProvider(usersvcProvider))
	usersvcClientProxy := client.NewUsersvcClientProxy(usersvcClient)

	svc := service.NewOrdersvc(conf, nil, usersvcClientProxy)

	handler := httpsrv.NewOrdersvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
