package client

import (
	"context"
	"ordersvc/vo"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/slok/goresilience"
	"github.com/slok/goresilience/circuitbreaker"
	rerrors "github.com/slok/goresilience/errors"
	"github.com/slok/goresilience/metrics"
	"github.com/slok/goresilience/retry"
	"github.com/slok/goresilience/timeout"
)

type OrdersvcClientProxy struct {
	client *OrdersvcClient
	logger *logrus.Logger
	runner goresilience.Runner
}

func (receiver *OrdersvcClientProxy) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, code, data, err = receiver.client.PageUsers(
			ctx,
			query,
		)
		if err != nil {
			return errors.Wrap(err, "call PageUsers fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call PageUsers fail")
	}
	return
}
func (receiver *OrdersvcClientProxy) GetHello(ctx context.Context) (ret string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, ret, err = receiver.client.GetHello(
			ctx,
		)
		if err != nil {
			return errors.Wrap(err, "call GetHello fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetHello fail")
	}
	return
}
func (receiver *OrdersvcClientProxy) GetGreeting(ctx context.Context, hello string) (ret string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, ret, err = receiver.client.GetGreeting(
			ctx,
			hello,
		)
		if err != nil {
			return errors.Wrap(err, "call GetGreeting fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetGreeting fail")
	}
	return
}
func (receiver *OrdersvcClientProxy) GetHelloWorld(ctx context.Context) (ret string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, ret, err = receiver.client.GetHelloWorld(
			ctx,
		)
		if err != nil {
			return errors.Wrap(err, "call GetHelloWorld fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetHelloWorld fail")
	}
	return
}

type ProxyOption func(*OrdersvcClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *OrdersvcClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger *logrus.Logger) ProxyOption {
	return func(proxy *OrdersvcClientProxy) {
		proxy.logger = logger
	}
}

func NewOrdersvcClientProxy(client *OrdersvcClient, opts ...ProxyOption) *OrdersvcClientProxy {
	cp := &OrdersvcClientProxy{
		client: client,
		logger: logrus.StandardLogger(),
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("ordersvc_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
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
