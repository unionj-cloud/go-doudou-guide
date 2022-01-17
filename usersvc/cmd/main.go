package main

import (
	"context"
	"fmt"
	"github.com/ascarter/requestid"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/ratelimit"
	"github.com/unionj-cloud/go-doudou/ratelimit/redisrate"
	ddconfig "github.com/unionj-cloud/go-doudou/svc/config"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
	"github.com/unionj-cloud/go-doudou/svc/logger"
	"github.com/unionj-cloud/go-doudou/svc/registry"
	"github.com/unionj-cloud/go-doudou/svc/tracing"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	service "usersvc"
	"usersvc/config"
	"usersvc/transport/httpsrv"
)

func main() {
	ddconfig.InitEnv()
	conf := config.LoadFromEnv()

	if logger.CheckDev() {
		logger.Init(logger.WithWritter(os.Stdout))
	} else {
		logger.Init(logger.WithWritter(io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename:   filepath.Join(os.Getenv("LOG_PATH"), fmt.Sprintf("%s.log", ddconfig.GddServiceName.Load())),
			MaxSize:    5,  // Max megabytes before log is rotated
			MaxBackups: 10, // Max number of old log files to keep
			MaxAge:     7,  // Max number of days to retain log files
			Compress:   true,
		})))
	}

	if ddconfig.GddMode.Load() == "micro" {
		err := registry.NewNode()
		if err != nil {
			logrus.Panicln(fmt.Sprintf("%+v", err))
		}
		defer registry.Shutdown()
	}

	tracer, closer := tracing.Init()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	svc := service.NewUsersvc(conf)

	handler := httpsrv.NewUsersvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()

	//store := memrate.NewMemoryStore(func(_ context.Context, store *memrate.MemoryStore, key string) ratelimit.Limiter {
	//	return memrate.NewLimiter(10, 30, memrate.WithTimer(10*time.Second, func() {
	//		store.DeleteKey(key)
	//	}))
	//})

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	fn := redisrate.LimitFn(func(ctx context.Context) ratelimit.Limit {
		return ratelimit.PerSecondBurst(10, 30)
	})

	srv.AddMiddleware(ddhttp.Tracing, ddhttp.Metrics,
		//ddhttp.BulkHead(1, 10*time.Millisecond),
		requestid.RequestIDHandler, handlers.CompressHandler, handlers.ProxyHeaders,
		//httpsrv.RateLimit(store),
		httpsrv.RedisRateLimit(rdb, fn),
		ddhttp.Logger,
		ddhttp.Rest, ddhttp.Recover)
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
