package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	ddconfig "github.com/unionj-cloud/go-doudou/svc/config"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
	"github.com/unionj-cloud/go-doudou/svc/registry"
)

func main() {
	ddconfig.InitEnv()
	node, err := registry.NewNode()
	if err != nil {
		logrus.Panicln(fmt.Sprintf("%+v", err))
	}
	logrus.Infof("Memberlist created. Local node is %s\n", node)

	srv := ddhttp.NewDefaultHttpSrv()
	srv.Run()
}
