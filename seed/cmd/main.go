package main

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/sirupsen/logrus"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	"github.com/unionj-cloud/go-doudou/toolkit/fileutils"
	"github.com/unionj-cloud/go-doudou/toolkit/stringutils"
	"os"
	"path/filepath"
	"seed/adapter"
	"seed/promsd"
	"time"
)

func main() {
	err := registry.NewNode()
	if err != nil {
		logrus.Panicln(fmt.Sprintf("%+v", err))
	}
	defer registry.Shutdown()

	interval, _ := time.ParseDuration(os.Getenv("PROM_REFRESH_INTERVAL"))
	if interval == 0 {
		interval = 30 * time.Second
	}

	kitLogger := log.NewSyncLogger(log.NewLogfmtLogger(os.Stdout))
	kitLogger = log.With(kitLogger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	out := os.Getenv("PROM_SD_OUT")
	if stringutils.IsNotEmpty(out) {
		_ = fileutils.CreateDirectory(out)
	}
	sdFile := filepath.Join(out, "go-doudou.json")

	disc, err := promsd.NewDiscovery(interval, kitLogger, sdFile)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sdAdapter := adapter.NewAdapter(ctx, sdFile, "go-doudou", disc, kitLogger)
	sdAdapter.Run()

	srv := ddhttp.NewDefaultHttpSrv()
	srv.Run()
}
