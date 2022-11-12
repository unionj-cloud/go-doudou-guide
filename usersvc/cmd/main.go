package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/unionj-cloud/go-doudou/v2/framework/ratelimit"
	"github.com/unionj-cloud/go-doudou/v2/framework/ratelimit/memrate"
	"github.com/unionj-cloud/go-doudou/v2/framework/ratelimit/redisrate"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/framework/tracing"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"time"
	service "usersvc"
	"usersvc/config"
	"usersvc/transport/httpsrv"
)

func main() {
	conf := config.LoadFromEnv()
	//if !framework.CheckDev() {
	zlogger.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   filepath.Join(os.Getenv("LOG_PATH"), fmt.Sprintf("%s.log", "usersvc")),
		MaxSize:    5,  // Max megabytes before log is rotated
		MaxBackups: 10, // Max number of old log files to keep
		MaxAge:     7,  // Max number of days to retain log files
		Compress:   true,
	}))
	//}
	tracer, closer := tracing.Init()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	svc := service.NewUsersvc(conf)

	handler := httpsrv.NewUsersvcHandler(svc)
	srv := rest.NewRestServer()

	store := memrate.NewMemoryStore(func(_ context.Context, store *memrate.MemoryStore, key string) ratelimit.Limiter {
		return memrate.NewLimiter(10, 30, memrate.WithTimer(10*time.Second, func() {
			store.DeleteKey(key)
		}))
	})

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	fn := redisrate.LimitFn(func(ctx context.Context) ratelimit.Limit {
		return ratelimit.PerSecondBurst(10, 30)
	})

	srv.AddMiddleware(
		rest.BulkHead(1, 10*time.Millisecond),
		httpsrv.RateLimit(store),
		httpsrv.RedisRateLimit(rdb, fn),
	)
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
