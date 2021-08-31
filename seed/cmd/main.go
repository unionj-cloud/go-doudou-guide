package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
	"github.com/unionj-cloud/go-doudou/svc/registry"
)

func main() {
	node, err := registry.NewNode()
	if err != nil {
		logrus.Panicln(fmt.Sprintf("%+v", err))
	}
	logrus.Infof("Memberlist created. Local node is %s\n", node)

	srv := ddhttp.NewDefaultHttpSrv()
	srv.Run()
}
