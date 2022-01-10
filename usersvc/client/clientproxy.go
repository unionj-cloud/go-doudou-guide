package client

import (
	"context"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/slok/goresilience"
	"github.com/slok/goresilience/circuitbreaker"
	rerrors "github.com/slok/goresilience/errors"
	"github.com/slok/goresilience/metrics"
	"github.com/slok/goresilience/retry"
	"github.com/slok/goresilience/timeout"
	v3 "github.com/unionj-cloud/go-doudou/openapi/v3"
	"github.com/unionj-cloud/go-doudou/svc/config"
	"os"
	"time"
	service "usersvc"
	"usersvc/vo"
)

type ClientProxy struct {
	client service.Usersvc
	logger *logrus.Logger
	runner goresilience.Runner
}

func (c ClientProxy) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error) {
	err := c.runner.Run(ctx, func(ctx context.Context) error {
		code, data, msg = c.client.PageUsers(ctx, query)
		if msg != nil {
			return errors.Wrap(msg, "")
		}
		return nil
	})
	if err != nil {
		if errors.Is(err, rerrors.ErrCircuitOpen) {
			// you can implement your fallback logic here
			c.logger.Error(err)
		}
		msg = errors.Wrap(err, "")
	}
	return
}

func (c ClientProxy) GetUser(ctx context.Context, userId string, photo string) (code int, data string, msg error) {
	//TODO implement me
	panic("implement me")
}

func (c ClientProxy) SignUp(ctx context.Context, username string, password int, actived bool, score float64) (code int, data string, msg error) {
	//TODO implement me
	panic("implement me")
}

func (c ClientProxy) UploadAvatar(ctx context.Context, models []*v3.FileModel, s string) (int, string, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClientProxy) UploadAvatar2(ctx context.Context, models []*v3.FileModel, s string, model *v3.FileModel, model2 *v3.FileModel) (int, string, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClientProxy) GetDownloadAvatar(ctx context.Context, userId string) (string, *os.File, error) {
	//TODO implement me
	panic("implement me")
}

type ProxyOption func(*ClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *ClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger *logrus.Logger) ProxyOption {
	return func(proxy *ClientProxy) {
		proxy.logger = logger
	}
}

func NewClientProxy(client service.Usersvc, opts ...ProxyOption) *ClientProxy {
	cp := &ClientProxy{
		client: client,
		logger: logrus.StandardLogger(),
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware

		if config.GddManage.Load() == "true" {
			mid = append(mid, metrics.NewMiddleware("usersvc_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
		}

		mid = append(mid, circuitbreaker.NewMiddleware(circuitbreaker.Config{
			ErrorPercentThresholdToOpen:        50,
			MinimumRequestToOpen:               6,
			SuccessfulRequiredOnHalfOpen:       1,
			WaitDurationInOpenState:            5 * time.Second,
			MetricsSlidingWindowBucketQuantity: 10,
			MetricsBucketDuration:              1 * time.Second,
		}),
			timeout.NewMiddleware(timeout.Config{
				Timeout: 3 * time.Minute,
			}),
			retry.NewMiddleware(retry.Config{
				Times: 3,
			}))

		cp.runner = goresilience.RunnerChain(mid...)
	}

	return cp
}
