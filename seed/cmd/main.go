package main

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/sirupsen/logrus"
	ddconfig "github.com/unionj-cloud/go-doudou/svc/config"
	"github.com/unionj-cloud/go-doudou/svc/logger"
	"github.com/unionj-cloud/go-doudou/svc/registry"
	"os"
	"seed/adapter"
	"seed/promsd"
	"time"
)

func main() {
	ddconfig.InitEnv()
	logger.Init()
	err := registry.NewNode()
	if err != nil {
		logrus.Panicln(fmt.Sprintf("%+v", err))
	}
	defer registry.Shutdown()

	ctx := context.Background()

	interval, _ := time.ParseDuration(os.Getenv("PROM_REFRESH_INTERVAL"))
	if interval == 0 {
		interval = 30 * time.Second
	}

	disc, err := promsd.NewDiscovery(interval)
	if err != nil {
		panic(err)
	}

	kitLogger := log.NewSyncLogger(log.NewLogfmtLogger(os.Stdout))
	kitLogger = log.With(kitLogger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	sdAdapter := adapter.NewAdapter(ctx, "go-doudou_sd.json", "goDoudouSD", disc, kitLogger)
	sdAdapter.Run()

	<-ctx.Done()
}
