package client

import (
	"context"
	"os"
	"time"
	"usersvc/vo"

	"github.com/go-resty/resty/v2"
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

	v3 "github.com/unionj-cloud/go-doudou/v2/toolkit/openapi/v3"
)

type UsersvcClientProxy struct {
	client *UsersvcClient
	logger *logrus.Logger
	runner goresilience.Runner
}

func (receiver *UsersvcClientProxy) PageUsers(ctx context.Context, _headers map[string]string, query vo.PageQuery) (_resp *resty.Response, code int, data vo.PageRet, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, code, data, msg = receiver.client.PageUsers(
			ctx,
			_headers,
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
func (receiver *UsersvcClientProxy) GetUser(ctx context.Context, _headers map[string]string, userId string, photo string) (_resp *resty.Response, code int, data string, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, code, data, msg = receiver.client.GetUser(
			ctx,
			_headers,
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
func (receiver *UsersvcClientProxy) SignUp(ctx context.Context, _headers map[string]string, username string, password int, actived bool, score float64) (_resp *resty.Response, code int, data string, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, code, data, msg = receiver.client.SignUp(
			ctx,
			_headers,
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
func (receiver *UsersvcClientProxy) UploadAvatar(ctx context.Context, _headers map[string]string, pf []v3.FileModel, ps string) (_resp *resty.Response, ri int, rs string, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, ri, rs, re = receiver.client.UploadAvatar(
			ctx,
			_headers,
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
func (receiver *UsersvcClientProxy) UploadAvatar2(ctx context.Context, _headers map[string]string, pf []v3.FileModel, ps string, pf2 *v3.FileModel, pf3 *v3.FileModel) (_resp *resty.Response, ri int, rs string, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, ri, rs, re = receiver.client.UploadAvatar2(
			ctx,
			_headers,
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
func (receiver *UsersvcClientProxy) GetDownloadAvatar(ctx context.Context, _headers map[string]string, userId string) (_resp *resty.Response, rs string, rf *os.File, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, rs, rf, re = receiver.client.GetDownloadAvatar(
			ctx,
			_headers,
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
func (receiver *UsersvcClientProxy) GetUser2(ctx context.Context, _headers map[string]string, userId string, photo *string) (_resp *resty.Response, code int, data *string, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, code, data, msg = receiver.client.GetUser2(
			ctx,
			_headers,
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
func (receiver *UsersvcClientProxy) PageUsers2(ctx context.Context, _headers map[string]string, query *vo.PageQuery) (_resp *resty.Response, code int, data vo.PageRet, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, code, data, msg = receiver.client.PageUsers2(
			ctx,
			_headers,
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
func (receiver *UsersvcClientProxy) GetUser3(ctx context.Context, _headers map[string]string, userId string, photo *string, attrs []int, pattrs *[]int) (_resp *resty.Response, code int, data *string, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, code, data, msg = receiver.client.GetUser3(
			ctx,
			_headers,
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
func (receiver *UsersvcClientProxy) GetUser4(ctx context.Context, _headers map[string]string, userId string, photo *string, pattrs *[]int, attrs2 ...int) (_resp *resty.Response, code int, data *string, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, code, data, msg = receiver.client.GetUser4(
			ctx,
			_headers,
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

func (receiver *UsersvcClientProxy) PageUsers3(ctx context.Context, _headers map[string]string, query vo.PageQuery1) (_resp *resty.Response, code int, data vo.PageRet, msg error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, code, data, msg = receiver.client.PageUsers3(
			ctx,
			_headers,
			query,
		)
		if msg != nil {
			return errors.Wrap(msg, "call PageUsers3 fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		msg = errors.Wrap(_err, "call PageUsers3 fail")
	}
	return
}
