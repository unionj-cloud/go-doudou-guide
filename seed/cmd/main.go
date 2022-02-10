package main

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/sirupsen/logrus"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	"github.com/unionj-cloud/go-doudou/toolkit/fileutils"
	"github.com/unionj-cloud/go-doudou/toolkit/stringutils"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"seed/adapter"
	"seed/promsd"
	"time"
)

func main() {
	if logger.CheckDev() {
		logger.Init(logger.WithWritter(os.Stdout))
	} else {
		logger.Init(logger.WithWritter(io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename:   filepath.Join(os.Getenv("LOG_PATH"), fmt.Sprintf("%s.log", "seed")),
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
