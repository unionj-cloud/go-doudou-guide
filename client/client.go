package client

import (
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/unionj-cloud/go-doudou/stringutils"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	service "usersvc"
	"usersvc/vo"
)

type IServiceProvider interface {
	SelectServer() (string, error)
}

type ServiceProvider struct {
}

func (s *ServiceProvider) SelectServer() (string, error) {
	address := os.Getenv("USERSVC")
	if stringutils.IsEmpty(address) {
		return "", errors.New("No service address for Usersvc found!")
	}
	return address, nil
}

func newServiceProvider() IServiceProvider {
	return &ServiceProvider{}
}

type UsersvcClient struct {
	provider IServiceProvider
	client   *resty.Client
}

func (u *UsersvcClient) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error) {
	server, err := u.provider.SelectServer()
	if err != nil {
		msg = errors.Wrap(err, "")
		return
	}
	resp, err := u.client.R().
		SetBody(query).
		Post(server + "/usersvc/pageusers")
	if err != nil {
		msg = errors.Wrap(err, "")
		return
	}
	if resp.IsError() {
		msg = errors.New(resp.String())
		return
	}
	var result struct {
		Code int        `json:"code"`
		Data vo.PageRet `json:"data"`
		Msg  string     `json:"msg"`
	}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		msg = errors.Wrap(err, "")
		return
	}
	if stringutils.IsNotEmpty(result.Msg) {
		msg = errors.New(result.Msg)
		return
	}
	return result.Code, result.Data, nil
}

func (u *UsersvcClient) GetUser(ctx context.Context, userId string, photo string) (code int, data string, msg error) {
	server, err := u.provider.SelectServer()
	if err != nil {
		msg = errors.Wrap(err, "")
		return
	}
	v := url.Values{}
	v.Set("userId", userId)
	v.Set("photo", photo)
	resp, err := u.client.R().
		SetQueryParamsFromValues(v).
		Get(server + "/usersvc/user")
	if err != nil {
		msg = errors.Wrap(err, "")
		return
	}
	if resp.IsError() {
		msg = errors.New(resp.String())
		return
	}
	var result struct {
		Code int    `json:"code"`
		Data string `json:"data"`
		Msg  string `json:"msg"`
	}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		msg = errors.Wrap(err, "")
		return
	}
	if stringutils.IsNotEmpty(result.Msg) {
		msg = errors.New(result.Msg)
		return
	}
	return result.Code, result.Data, nil
}

func (u *UsersvcClient) SignUp(ctx context.Context, username string, password string) (code int, data string, msg error) {
	server, err := u.provider.SelectServer()
	if err != nil {
		msg = errors.Wrap(err, "")
		return
	}
	v := url.Values{}
	v.Set("username", username)
	v.Set("password", password)
	resp, err := u.client.R().
		SetFormDataFromValues(v).
		Post(server + "/usersvc/signup")
	if err != nil {
		msg = errors.Wrap(err, "")
		return
	}
	if resp.IsError() {
		msg = errors.New(resp.String())
		return
	}
	var result struct {
		Code int    `json:"code"`
		Data string `json:"data"`
		Msg  string `json:"msg"`
	}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		msg = errors.Wrap(err, "")
		return
	}
	if stringutils.IsNotEmpty(result.Msg) {
		msg = errors.New(result.Msg)
		return
	}
	return result.Code, result.Data, nil
}

func (u *UsersvcClient) UploadAvatar(ctx context.Context, headers []*multipart.FileHeader, s string) (int, string, error) {
	server, err := u.provider.SelectServer()
	if err != nil {
		err = errors.Wrap(err, "")
		return 0, "", err
	}
	req := u.client.R()
	for _, fh := range headers {
		f, err := fh.Open()
		if err != nil {
			err = errors.Wrap(err, "")
			return 0, "", err
		}
		req.SetFileReader("headers", fh.Filename, f)
	}
	v := url.Values{}
	v.Set("s", s)
	resp, err := req.SetFormDataFromValues(v).Post(server + "/usersvc/uploadavatar")
	if err != nil {
		err = errors.Wrap(err, "")
		return 0, "", err
	}
	if resp.IsError() {
		err = errors.New(resp.String())
		return 0, "", err
	}
	var result struct {
		I int    `json:"i"`
		S string `json:"s"`
		E string `json:"e"`
	}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		err = errors.Wrap(err, "")
		return 0, "", err
	}
	if stringutils.IsNotEmpty(result.E) {
		err = errors.New(result.E)
		return 0, "", err
	}
	return result.I, result.S, nil
}

func createDirectory(dir string) (err error) {
	if _, err = os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(dir, 0755); err != nil {
				return
			}
		}
	}
	return
}

func (u *UsersvcClient) DownloadAvatar(ctx context.Context, userId string) (*os.File, error) {
	server, err := u.provider.SelectServer()
	if err != nil {
		err = errors.Wrap(err, "")
		return nil, err
	}
	v := url.Values{}
	v.Set("userId", userId)
	resp, err := u.client.R().
		SetQueryParamsFromValues(v).
		SetDoNotParseResponse(true).
		Get(server + "/usersvc/downloadavatar")
	if err != nil {
		err = errors.Wrap(err, "")
		return nil, err
	}
	if resp.IsError() {
		err = errors.New(resp.String())
		return nil, err
	}
	disp := resp.Header().Get("Content-Disposition")
	file := strings.TrimPrefix(disp, "attachment; filename=")
	output := os.Getenv("OUTPUT")
	if stringutils.IsNotEmpty(output) {
		file = output + string(filepath.Separator) + file
	}
	file = filepath.Clean(file)
	if err = createDirectory(filepath.Dir(file)); err != nil {
		err = errors.Wrap(err, "")
		return nil, err
	}
	outFile, err := os.Create(file)
	if err != nil {
		err = errors.Wrap(err, "")
		return nil, err
	}
	defer outFile.Close()
	defer resp.RawBody().Close()
	_, err = io.Copy(outFile, resp.RawBody())
	if err != nil {
		err = errors.Wrap(err, "")
		return nil, err
	}
	return outFile, nil
}

type UsersvcClientOption func(*UsersvcClient)

func WithProvider(provider IServiceProvider) UsersvcClientOption {
	return func(a *UsersvcClient) {
		a.provider = provider
	}
}

func WithClient(client *resty.Client) UsersvcClientOption {
	return func(a *UsersvcClient) {
		a.client = client
	}
}

func newClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(1 * time.Minute)

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	client.SetTransport(&http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		MaxConnsPerHost:       100,
	})
	return client
}

func NewUsersvc(opts ...UsersvcClientOption) service.Usersvc {
	defaultProvider := newServiceProvider()
	defaultClient := newClient()

	a := &UsersvcClient{
		provider: defaultProvider,
		client:   defaultClient,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}
