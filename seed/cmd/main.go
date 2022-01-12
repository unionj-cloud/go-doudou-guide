package main

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/fileutils"
	"github.com/unionj-cloud/go-doudou/stringutils"
	ddconfig "github.com/unionj-cloud/go-doudou/svc/config"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
	"github.com/unionj-cloud/go-doudou/svc/logger"
	"github.com/unionj-cloud/go-doudou/svc/registry"
	"os"
	"path/filepath"
	"seed/adapter"
	"seed/promsd"
	"time"
)

func main() {
	ddconfig.InitEnv()
	logger.Init()
	logFile := logger.PersistLogToDisk()
	defer logFile.Close()

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

	disc, err := promsd.NewDiscovery(interval, kitLogger)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out := os.Getenv("PROM_SD_OUT")
	if stringutils.IsNotEmpty(out) {
		_ = fileutils.CreateDirectory(out)
	}

	sdAdapter := adapter.NewAdapter(ctx, filepath.Join(out, "godoudouguide_sd.json"), "godoudouguideSD", disc, kitLogger)
	sdAdapter.Run()

	srv := ddhttp.NewDefaultHttpSrv()
	srv.Run()
}
