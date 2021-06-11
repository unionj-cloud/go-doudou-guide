package client

import (
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	service "usersvc"
	"usersvc/vo"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/fileutils"
	"github.com/unionj-cloud/go-doudou/stringutils"
	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
)

type UsersvcClient struct {
	provider ddhttp.IServiceProvider
	client   *resty.Client
}

func (receiver *UsersvcClient) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error) {
	var (
		_server string
		_err    error
	)
	if _server, _err = receiver.provider.SelectServer(); _err != nil {
		msg = errors.Wrap(_err, "")
		return
	}
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetBody(query)
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err := _req.Post(_server + "/usersvc/pageusers")
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
	var (
		_server string
		_err    error
	)
	if _server, _err = receiver.provider.SelectServer(); _err != nil {
		msg = errors.Wrap(_err, "")
		return
	}
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_urlValues.Set("userId", userId)
	_urlValues.Set("photo", photo)
	_resp, _err := _req.SetQueryParamsFromValues(_urlValues).
		Get(_server + "/usersvc/user")
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
func (receiver *UsersvcClient) SignUp(ctx context.Context, username string, password string) (code int, data string, msg error) {
	var (
		_server string
		_err    error
	)
	if _server, _err = receiver.provider.SelectServer(); _err != nil {
		msg = errors.Wrap(_err, "")
		return
	}
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_urlValues.Set("username", username)
	_urlValues.Set("password", password)
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err := _req.Post(_server + "/usersvc/signup")
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
func (receiver *UsersvcClient) UploadAvatar(pc context.Context, pf []*multipart.FileHeader, ps string) (ri int, rs string, re error) {
	var (
		_server string
		_err    error
	)
	if _server, _err = receiver.provider.SelectServer(); _err != nil {
		re = errors.Wrap(_err, "")
		return
	}
	_urlValues := url.Values{}
	_req := receiver.client.R()
	for _, _fh := range pf {
		_f, _err := _fh.Open()
		if _err != nil {
			re = errors.Wrap(_err, "")
			return
		}
		_req.SetFileReader("pf", _fh.Filename, _f)
	}
	_urlValues.Set("ps", ps)
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err := _req.Post(_server + "/usersvc/uploadavatar")
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
func (receiver *UsersvcClient) DownloadAvatar(ctx context.Context, userId string) (rf *os.File, re error) {
	var (
		_server string
		_err    error
	)
	if _server, _err = receiver.provider.SelectServer(); _err != nil {
		re = errors.Wrap(_err, "")
		return
	}
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_urlValues.Set("userId", userId)
	_req.SetDoNotParseResponse(true)
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err := _req.Post(_server + "/usersvc/downloadavatar")
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
	_output := os.Getenv("OUTPUT")
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

type UsersvcClientOption func(*UsersvcClient)

func WithProvider(provider ddhttp.IServiceProvider) UsersvcClientOption {
	return func(a *UsersvcClient) {
		a.provider = provider
	}
}

func WithClient(client *resty.Client) UsersvcClientOption {
	return func(a *UsersvcClient) {
		a.client = client
	}
}

func NewUsersvc(opts ...UsersvcClientOption) service.Usersvc {
	defaultProvider := ddhttp.NewServiceProvider("Usersvc")
	defaultClient := ddhttp.NewClient()

	svcClient := &UsersvcClient{
		provider: defaultProvider,
		client:   defaultClient,
	}

	for _, opt := range opts {
		opt(svcClient)
	}

	return svcClient
}
