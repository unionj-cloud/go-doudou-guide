package client

import (
	"context"
	"os"
	"time"
	"usersvc/vo"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/slok/goresilience"
	"github.com/slok/goresilience/circuitbreaker"
	rerrors "github.com/slok/goresilience/errors"
	"github.com/slok/goresilience/metrics"
	"github.com/slok/goresilience/retry"
	"github.com/slok/goresilience/timeout"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type UsersvcClientProxy struct {
	client *UsersvcClient
	logger *logrus.Logger
	runner goresilience.Runner
}

func (receiver *UsersvcClientProxy) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, code, data, msg = receiver.client.PageUsers(
			ctx,
			query,
		)
		if msg != nil {
			return errors.Wrap(msg, "call PageUsers fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		msg = errors.Wrap(_err, "call PageUsers fail")
	}
	return
}
func (receiver *UsersvcClientProxy) GetUser(ctx context.Context, userId string, photo string) (code int, data string, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, code, data, msg = receiver.client.GetUser(
			ctx,
			userId,
			photo,
		)
		if msg != nil {
			return errors.Wrap(msg, "call GetUser fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		msg = errors.Wrap(_err, "call GetUser fail")
	}
	return
}
func (receiver *UsersvcClientProxy) SignUp(ctx context.Context, username string, password int, actived bool, score float64) (code int, data string, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, code, data, msg = receiver.client.SignUp(
			ctx,
			username,
			password,
			actived,
			score,
		)
		if msg != nil {
			return errors.Wrap(msg, "call SignUp fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		msg = errors.Wrap(_err, "call SignUp fail")
	}
	return
}
func (receiver *UsersvcClientProxy) UploadAvatar(pc context.Context, pf []v3.FileModel, ps string) (ri int, rs string, re error) {
	if _err := receiver.runner.Run(pc, func(ctx context.Context) error {
		_, ri, rs, re = receiver.client.UploadAvatar(
			pc,
			pf,
			ps,
		)
		if re != nil {
			return errors.Wrap(re, "call UploadAvatar fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		re = errors.Wrap(_err, "call UploadAvatar fail")
	}
	return
}
func (receiver *UsersvcClientProxy) UploadAvatar2(pc context.Context, pf []v3.FileModel, ps string, pf2 *v3.FileModel, pf3 *v3.FileModel) (ri int, rs string, re error) {
	if _err := receiver.runner.Run(pc, func(ctx context.Context) error {
		_, ri, rs, re = receiver.client.UploadAvatar2(
			pc,
			pf,
			ps,
			pf2,
			pf3,
		)
		if re != nil {
			return errors.Wrap(re, "call UploadAvatar2 fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		re = errors.Wrap(_err, "call UploadAvatar2 fail")
	}
	return
}
func (receiver *UsersvcClientProxy) GetDownloadAvatar(ctx context.Context, userId string) (rs string, rf *os.File, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, rs, rf, re = receiver.client.GetDownloadAvatar(
			ctx,
			userId,
		)
		if re != nil {
			return errors.Wrap(re, "call GetDownloadAvatar fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		re = errors.Wrap(_err, "call GetDownloadAvatar fail")
	}
	return
}
func (receiver *UsersvcClientProxy) GetUser2(ctx context.Context, userId string, photo *string) (code int, data *string, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, code, data, msg = receiver.client.GetUser2(
			ctx,
			userId,
			photo,
		)
		if msg != nil {
			return errors.Wrap(msg, "call GetUser2 fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		msg = errors.Wrap(_err, "call GetUser2 fail")
	}
	return
}
func (receiver *UsersvcClientProxy) PageUsers2(ctx context.Context, query *vo.PageQuery) (code int, data vo.PageRet, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, code, data, msg = receiver.client.PageUsers2(
			ctx,
			query,
		)
		if msg != nil {
			return errors.Wrap(msg, "call PageUsers2 fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		msg = errors.Wrap(_err, "call PageUsers2 fail")
	}
	return
}
func (receiver *UsersvcClientProxy) GetUser3(ctx context.Context, userId string, photo *string, attrs []int, pattrs *[]int) (code int, data *string, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, code, data, msg = receiver.client.GetUser3(
			ctx,
			userId,
			photo,
			attrs,
			pattrs,
		)
		if msg != nil {
			return errors.Wrap(msg, "call GetUser3 fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		msg = errors.Wrap(_err, "call GetUser3 fail")
	}
	return
}
func (receiver *UsersvcClientProxy) GetUser4(ctx context.Context, userId string, photo *string, pattrs *[]int, attrs2 ...int) (code int, data *string, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, code, data, msg = receiver.client.GetUser4(
			ctx,
			userId,
			photo,
			pattrs,
			attrs2...,
		)
		if msg != nil {
			return errors.Wrap(msg, "call GetUser4 fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		msg = errors.Wrap(_err, "call GetUser4 fail")
	}
	return
}

type ProxyOption func(*UsersvcClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *UsersvcClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger *logrus.Logger) ProxyOption {
	return func(proxy *UsersvcClientProxy) {
		proxy.logger = logger
	}
}

func NewUsersvcClientProxy(client *UsersvcClient, opts ...ProxyOption) *UsersvcClientProxy {
	cp := &UsersvcClientProxy{
		client: client,
		logger: logrus.StandardLogger(),
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("usersvc_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
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
