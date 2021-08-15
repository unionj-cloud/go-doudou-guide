package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/pathutils"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
	"github.com/unionj-cloud/go-doudou/svc/registry"
)

func main() {
	err := godotenv.Load(pathutils.Abs("../.env"))
	if err != nil {
		logrus.Fatal("Error loading .env file", err)
	}

	node, err := registry.NewNode()
	if err != nil {
		logrus.Panicln(fmt.Sprintf("%+v", err))
	}
	logrus.Infof("Memberlist created. Local node is %s\n", node)

	//ips, err := net.LookupIP("seed-2.seed-svc-headless.default.svc.cluster.local")
	//if err != nil {
	//	logrus.Error(err)
	//}
	//for _, ip := range ips {
	//	logrus.Infoln(ip.String())
	//}

	srv := ddhttp.NewDefaultHttpSrv()
	srv.Run()
}
