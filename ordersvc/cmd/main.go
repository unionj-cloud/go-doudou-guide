package main

import (
	"fmt"
	"github.com/ascarter/requestid"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
	"github.com/unionj-cloud/go-doudou/svc/registry"
	service "ordersvc"
	"ordersvc/config"
	"ordersvc/transport/httpsrv"
	"usersvc/client"
)

func main() {
	conf := config.LoadFromEnv()
	//conn, err := db.NewDb(conf.DbConf)
	//if err != nil {
	//	panic(err)
	//}
	//defer func() {
	//	if conn == nil {
	//		return
	//	}
	//	if err := conn.Close(); err == nil {
	//		logrus.Infoln("Database connection is closed")
	//	} else {
	//		logrus.Warnln("Failed to close database connection")
	//	}
	//}()

	node, err := registry.NewNode()
	if err != nil {
		logrus.Panicln(fmt.Sprintf("%+v", err))
	}
	logrus.Infof("%s joined cluster\n", node.String())

	usersvcProvider := ddhttp.NewMemberlistServiceProvider("usersvc", node)
	usersvcClient := client.NewUsersvc(ddhttp.WithProvider(usersvcProvider))

	svc := service.NewOrdersvc(conf, nil, usersvcClient)

	handler := httpsrv.NewOrdersvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddMiddleware(ddhttp.Metrics, requestid.RequestIDHandler, handlers.CompressHandler, handlers.ProxyHeaders, ddhttp.Logger, ddhttp.Rest)
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
