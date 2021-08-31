package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"ordersvc/vo"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/stringutils"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
)

type OrdersvcClient struct {
	provider ddhttp.IServiceProvider
	client   *resty.Client
}

func (receiver *OrdersvcClient) SetProvider(provider ddhttp.IServiceProvider) {
	receiver.provider = provider
}

func (receiver *OrdersvcClient) SetClient(client *resty.Client) {
	receiver.client = client
}
func (receiver *OrdersvcClient) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, err error) {
	var (
		_server string
		_err    error
	)
	if _server, _err = receiver.provider.SelectServer(); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_req.SetBody(query)
	_path := "/page/users"
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err := _req.Post(_server + _path)
	if _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int        `json:"code"`
		Data vo.PageRet `json:"data"`
		Err  string     `json:"err"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if stringutils.IsNotEmpty(_result.Err) {
		err = errors.New(_result.Err)
		return
	}
	return _result.Code, _result.Data, nil
}
func (receiver *OrdersvcClient) GetHello(ctx context.Context) (ret string, err error) {
	var (
		_server string
		_err    error
	)
	if _server, _err = receiver.provider.SelectServer(); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_path := "/hello"
	_resp, _err := _req.SetQueryParamsFromValues(_urlValues).
		Get(_server + _path)
	if _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		Ret string `json:"ret"`
		Err string `json:"err"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if stringutils.IsNotEmpty(_result.Err) {
		err = errors.New(_result.Err)
		return
	}
	return _result.Ret, nil
}
func (receiver *OrdersvcClient) GetGreeting(ctx context.Context, hello string) (ret string, err error) {
	var (
		_server string
		_err    error
	)
	if _server, _err = receiver.provider.SelectServer(); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("hello", fmt.Sprintf("%v", hello))
	_path := "/greeting"
	_resp, _err := _req.SetQueryParamsFromValues(_urlValues).
		Get(_server + _path)
	if _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		Ret string `json:"ret"`
		Err string `json:"err"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if stringutils.IsNotEmpty(_result.Err) {
		err = errors.New(_result.Err)
		return
	}
	return _result.Ret, nil
}
func (receiver *OrdersvcClient) GetHelloWorld(ctx context.Context) (ret string, err error) {
	var (
		_server string
		_err    error
	)
	if _server, _err = receiver.provider.SelectServer(); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_path := "/hello/world"
	_resp, _err := _req.SetQueryParamsFromValues(_urlValues).
		Get(_server + _path)
	if _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		Ret string `json:"ret"`
		Err string `json:"err"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if stringutils.IsNotEmpty(_result.Err) {
		err = errors.New(_result.Err)
		return
	}
	return _result.Ret, nil
}

func NewOrdersvc(opts ...ddhttp.DdClientOption) *OrdersvcClient {
	defaultProvider := ddhttp.NewServiceProvider("ORDERSVC")
	defaultClient := ddhttp.NewClient()

	svcClient := &OrdersvcClient{
		provider: defaultProvider,
		client:   defaultClient,
	}

	for _, opt := range opts {
		opt(svcClient)
	}

	return svcClient
}
