package main

import (
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/pathutils"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
	service "usersvc"
    "usersvc/config"
	"usersvc/db"
	"usersvc/transport/httpsrv"
)

func main() {
	env := config.NewDotenv(pathutils.Abs("../.env"))
	conf := env.Get()

	conn, err := db.NewDb(conf.DbConf)
	if err != nil {
		panic(err)
	}
	defer func() {
		if conn == nil {
			return
		}
		if err := conn.Close(); err == nil {
			logrus.Infoln("Database connection is closed")
		} else {
			logrus.Warnln("Failed to close database connection")
		}
	}()

    svc := service.NewUsersvc(conf, conn)

	handler := httpsrv.NewUsersvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddMiddleware(httpsrv.Logger, httpsrv.Rest)
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
