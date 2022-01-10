package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"usersvc/vo"

	"github.com/go-resty/resty/v2"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/fileutils"
	v3 "github.com/unionj-cloud/go-doudou/openapi/v3"
	"github.com/unionj-cloud/go-doudou/stringutils"
	"github.com/unionj-cloud/go-doudou/svc/config"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
	"github.com/unionj-cloud/go-doudou/svc/registry"
)

type UsersvcClient struct {
	provider registry.IServiceProvider
	client   *resty.Client
}

func (receiver *UsersvcClient) SetProvider(provider registry.IServiceProvider) {
	receiver.provider = provider
}

func (receiver *UsersvcClient) SetClient(client *resty.Client) {
	receiver.client = client
}
func (receiver *UsersvcClient) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error) {
	var _err error
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
	_resp, _err := _req.Post(_path)
	if _err != nil {
		msg = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		msg = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int        `json:"code"`
		Data vo.PageRet `json:"data"`
		Msg  string     `json:"msg"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		msg = errors.Wrap(_err, "")
		return
	}
	if stringutils.IsNotEmpty(_result.Msg) {
		msg = errors.New(_result.Msg)
		return
	}
	return _result.Code, _result.Data, nil
}
func (receiver *UsersvcClient) GetUser(ctx context.Context, userId string, photo string) (code int, data string, msg error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("userId", fmt.Sprintf("%v", userId))
	_urlValues.Set("photo", fmt.Sprintf("%v", photo))
	_path := "/user"
	_resp, _err := _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
	if _err != nil {
		msg = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		msg = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int    `json:"code"`
		Data string `json:"data"`
		Msg  string `json:"msg"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		msg = errors.Wrap(_err, "")
		return
	}
	if stringutils.IsNotEmpty(_result.Msg) {
		msg = errors.New(_result.Msg)
		return
	}
	return _result.Code, _result.Data, nil
}
func (receiver *UsersvcClient) SignUp(ctx context.Context, username string, password int, actived bool, score float64) (code int, data string, msg error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("username", fmt.Sprintf("%v", username))
	_urlValues.Set("password", fmt.Sprintf("%v", password))
	_urlValues.Set("actived", fmt.Sprintf("%v", actived))
	_urlValues.Set("score", fmt.Sprintf("%v", score))
	_path := "/sign/up"
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err := _req.Post(_path)
	if _err != nil {
		msg = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		msg = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int    `json:"code"`
		Data string `json:"data"`
		Msg  string `json:"msg"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		msg = errors.Wrap(_err, "")
		return
	}
	if stringutils.IsNotEmpty(_result.Msg) {
		msg = errors.New(_result.Msg)
		return
	}
	return _result.Code, _result.Data, nil
}
func (receiver *UsersvcClient) UploadAvatar(pc context.Context, pf []*v3.FileModel, ps string) (ri int, rs string, re error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(pc)
	for _, _f := range pf {
		_req.SetFileReader("pf", _f.Filename, _f.Reader)
	}
	_urlValues.Set("ps", fmt.Sprintf("%v", ps))
	_path := "/upload/avatar"
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err := _req.Post(_path)
	if _err != nil {
		re = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		re = errors.New(_resp.String())
		return
	}
	var _result struct {
		Ri int    `json:"ri"`
		Rs string `json:"rs"`
		Re string `json:"re"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		re = errors.Wrap(_err, "")
		return
	}
	if stringutils.IsNotEmpty(_result.Re) {
		re = errors.New(_result.Re)
		return
	}
	return _result.Ri, _result.Rs, nil
}
func (receiver *UsersvcClient) UploadAvatar2(pc context.Context, pf []*v3.FileModel, ps string, pf2 *v3.FileModel, pf3 *v3.FileModel) (ri int, rs string, re error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(pc)
	for _, _f := range pf {
		_req.SetFileReader("pf", _f.Filename, _f.Reader)
	}
	_urlValues.Set("ps", fmt.Sprintf("%v", ps))
	_req.SetFileReader("pf2", pf2.Filename, pf2.Reader)
	_req.SetFileReader("pf3", pf3.Filename, pf3.Reader)
	_path := "/upload/avatar/2"
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err := _req.Post(_path)
	if _err != nil {
		re = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		re = errors.New(_resp.String())
		return
	}
	var _result struct {
		Ri int    `json:"ri"`
		Rs string `json:"rs"`
		Re string `json:"re"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		re = errors.Wrap(_err, "")
		return
	}
	if stringutils.IsNotEmpty(_result.Re) {
		re = errors.New(_result.Re)
		return
	}
	return _result.Ri, _result.Rs, nil
}
func (receiver *UsersvcClient) GetDownloadAvatar(ctx context.Context, userId string) (rs string, rf *os.File, re error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("userId", fmt.Sprintf("%v", userId))
	_req.SetDoNotParseResponse(true)
	_path := "/download/avatar"
	_resp, _err := _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
	if _err != nil {
		re = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		re = errors.New(_resp.String())
		return
	}
	_disp := _resp.Header().Get("Content-Disposition")
	_file := strings.TrimPrefix(_disp, "attachment; filename=")
	_output := config.GddOutput.Load()
	if stringutils.IsNotEmpty(_output) {
		_file = _output + string(filepath.Separator) + _file
	}
	_file = filepath.Clean(_file)
	if _err = fileutils.CreateDirectory(filepath.Dir(_file)); _err != nil {
		re = errors.Wrap(_err, "")
		return
	}
	_outFile, _err := os.Create(_file)
	if _err != nil {
		re = errors.Wrap(_err, "")
		return
	}
	defer _outFile.Close()
	defer _resp.RawBody().Close()
	_, _err = io.Copy(_outFile, _resp.RawBody())
	if _err != nil {
		re = errors.Wrap(_err, "")
		return
	}
	rf = _outFile
	return
}

func NewUsersvc(opts ...ddhttp.DdClientOption) *UsersvcClient {
	defaultProvider := ddhttp.NewServiceProvider("USERSVC")
	defaultClient := ddhttp.NewClient()

	svcClient := &UsersvcClient{
		provider: defaultProvider,
		client:   defaultClient,
	}

	for _, opt := range opts {
		opt(svcClient)
	}

	svcClient.client.OnBeforeRequest(func(_ *resty.Client, request *resty.Request) error {
		request.URL = svcClient.provider.SelectServer() + request.URL
		return nil
	})

	svcClient.client.SetPreRequestHook(func(_ *resty.Client, request *http.Request) error {
		traceReq, _ := nethttp.TraceRequest(opentracing.GlobalTracer(), request,
			nethttp.OperationName(fmt.Sprintf("HTTP %s: %s", request.Method, request.RequestURI)))
		*request = *traceReq
		return nil
	})

	svcClient.client.OnAfterResponse(func(_ *resty.Client, response *resty.Response) error {
		nethttp.TracerFromRequest(response.Request.RawRequest).Finish()
		return nil
	})

	return svcClient
}
