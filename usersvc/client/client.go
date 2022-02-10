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
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	"github.com/unionj-cloud/go-doudou/toolkit/fileutils"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
	"github.com/unionj-cloud/go-doudou/toolkit/stringutils"
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
func (receiver *UsersvcClient) PageUsers(ctx context.Context, query vo.PageQuery) (_resp *resty.Response, code int, data vo.PageRet, msg error) {
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
	_resp, _err = _req.Post(_path)
	if _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		msg = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int        `json:"code"`
		Data vo.PageRet `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Code, _result.Data, nil
}
func (receiver *UsersvcClient) GetUser(ctx context.Context, userId string, photo string) (_resp *resty.Response, code int, data string, msg error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("userId", fmt.Sprintf("%v", userId))
	_urlValues.Set("photo", fmt.Sprintf("%v", photo))
	_path := "/user"
	_resp, _err = _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
	if _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		msg = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int    `json:"code"`
		Data string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Code, _result.Data, nil
}
func (receiver *UsersvcClient) SignUp(ctx context.Context, username string, password int, actived bool, score float64) (_resp *resty.Response, code int, data string, msg error) {
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
	_resp, _err = _req.Post(_path)
	if _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		msg = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int    `json:"code"`
		Data string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Code, _result.Data, nil
}
func (receiver *UsersvcClient) UploadAvatar(pc context.Context, pf []v3.FileModel, ps string) (_resp *resty.Response, ri int, rs string, re error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(pc)
	if len(pf) == 0 {
		re = errors.New("at least one file should be uploaded for parameter pf")
		return
	}
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
	_resp, _err = _req.Post(_path)
	if _err != nil {
		re = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		re = errors.New(_resp.String())
		return
	}
	var _result struct {
		Ri int    `json:"ri"`
		Rs string `json:"rs"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		re = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Ri, _result.Rs, nil
}
func (receiver *UsersvcClient) UploadAvatar2(pc context.Context, pf []v3.FileModel, ps string, pf2 *v3.FileModel, pf3 *v3.FileModel) (_resp *resty.Response, ri int, rs string, re error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(pc)
	if len(pf) == 0 {
		re = errors.New("at least one file should be uploaded for parameter pf")
		return
	}
	for _, _f := range pf {
		_req.SetFileReader("pf", _f.Filename, _f.Reader)
	}
	_urlValues.Set("ps", fmt.Sprintf("%v", ps))
	if pf2 != nil {
		_req.SetFileReader("pf2", pf2.Filename, pf2.Reader)
	}
	if pf3 != nil {
		_req.SetFileReader("pf3", pf3.Filename, pf3.Reader)
	}
	_path := "/upload/avatar/2"
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err = _req.Post(_path)
	if _err != nil {
		re = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		re = errors.New(_resp.String())
		return
	}
	var _result struct {
		Ri int    `json:"ri"`
		Rs string `json:"rs"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		re = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Ri, _result.Rs, nil
}
func (receiver *UsersvcClient) GetDownloadAvatar(ctx context.Context, userId string) (_resp *resty.Response, rs string, rf *os.File, re error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("userId", fmt.Sprintf("%v", userId))
	_req.SetDoNotParseResponse(true)
	_path := "/download/avatar"
	_resp, _err = _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
	if _err != nil {
		re = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		re = errors.New(_resp.String())
		return
	}
	_disp := _resp.Header().Get("Content-Disposition")
	_file := strings.TrimPrefix(_disp, "attachment; filename=")
	_output := os.TempDir()
	if stringutils.IsNotEmpty(_output) {
		_file = _output + string(filepath.Separator) + _file
	}
	_file = filepath.Clean(_file)
	if _err = fileutils.CreateDirectory(filepath.Dir(_file)); _err != nil {
		re = errors.Wrap(_err, "error")
		return
	}
	_outFile, _err := os.Create(_file)
	if _err != nil {
		re = errors.Wrap(_err, "error")
		return
	}
	defer _outFile.Close()
	defer _resp.RawBody().Close()
	_, _err = io.Copy(_outFile, _resp.RawBody())
	if _err != nil {
		re = errors.Wrap(_err, "error")
		return
	}
	rf = _outFile
	return
}
func (receiver *UsersvcClient) GetUser2(ctx context.Context, userId string, photo *string) (_resp *resty.Response, code int, data *string, msg error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("userId", fmt.Sprintf("%v", userId))
	if photo != nil {
		_urlValues.Set("photo", fmt.Sprintf("%v", *photo))
	}
	_path := "/user/2"
	_resp, _err = _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
	if _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		msg = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int     `json:"code"`
		Data *string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Code, _result.Data, nil
}
func (receiver *UsersvcClient) PageUsers2(ctx context.Context, query *vo.PageQuery) (_resp *resty.Response, code int, data vo.PageRet, msg error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_req.SetBody(query)
	_path := "/page/users/2"
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err = _req.Post(_path)
	if _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		msg = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int        `json:"code"`
		Data vo.PageRet `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Code, _result.Data, nil
}
func (receiver *UsersvcClient) GetUser3(ctx context.Context, userId string, photo *string, attrs []int, pattrs *[]int) (_resp *resty.Response, code int, data *string, msg error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("userId", fmt.Sprintf("%v", userId))
	if photo != nil {
		_urlValues.Set("photo", fmt.Sprintf("%v", *photo))
	}
	if len(attrs) == 0 {
		msg = errors.New("size of parameter attrs should be greater than zero")
		return
	}
	for _, _item := range attrs {
		_urlValues.Add("attrs", fmt.Sprintf("%v", _item))
	}
	if pattrs != nil {
		for _, _item := range *pattrs {
			_urlValues.Add("pattrs", fmt.Sprintf("%v", _item))
		}
	}
	_path := "/user/3"
	_resp, _err = _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
	if _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		msg = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int     `json:"code"`
		Data *string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Code, _result.Data, nil
}
func (receiver *UsersvcClient) GetUser4(ctx context.Context, userId string, photo *string, pattrs *[]int, attrs2 ...int) (_resp *resty.Response, code int, data *string, msg error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("userId", fmt.Sprintf("%v", userId))
	if photo != nil {
		_urlValues.Set("photo", fmt.Sprintf("%v", *photo))
	}
	if pattrs != nil {
		for _, _item := range *pattrs {
			_urlValues.Add("pattrs", fmt.Sprintf("%v", _item))
		}
	}
	if attrs2 != nil {
		for _, _item := range attrs2 {
			_urlValues.Add("attrs2", fmt.Sprintf("%v", _item))
		}
	}
	_path := "/user/4"
	_resp, _err = _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
	if _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		msg = errors.New(_resp.String())
		return
	}
	var _result struct {
		Code int     `json:"code"`
		Data *string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		msg = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Code, _result.Data, nil
}

func NewUsersvcClient(opts ...ddhttp.DdClientOption) *UsersvcClient {
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
