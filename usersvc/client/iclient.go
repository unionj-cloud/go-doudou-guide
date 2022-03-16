package client

import (
	"context"
	"os"
	"usersvc/vo"

	"github.com/go-resty/resty/v2"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type IUsersvcClient interface {
	PageUsers(ctx context.Context, _headers map[string]string, query vo.PageQuery) (_resp *resty.Response, code int, data vo.PageRet, msg error)
	GetUser(ctx context.Context, _headers map[string]string, userId string, photo string) (_resp *resty.Response, code int, data string, msg error)
	SignUp(ctx context.Context, _headers map[string]string, username string, password int, actived bool, score float64) (_resp *resty.Response, code int, data string, msg error)
	UploadAvatar(ctx context.Context, _headers map[string]string, pf []v3.FileModel, ps string) (_resp *resty.Response, ri int, rs string, re error)
	UploadAvatar2(ctx context.Context, _headers map[string]string, pf []v3.FileModel, ps string, pf2 *v3.FileModel, pf3 *v3.FileModel) (_resp *resty.Response, ri int, rs string, re error)
	GetDownloadAvatar(ctx context.Context, _headers map[string]string, userId string) (_resp *resty.Response, rs string, rf *os.File, re error)
	GetUser2(ctx context.Context, _headers map[string]string, userId string, photo *string) (_resp *resty.Response, code int, data *string, msg error)
	PageUsers2(ctx context.Context, _headers map[string]string, query *vo.PageQuery) (_resp *resty.Response, code int, data vo.PageRet, msg error)
	GetUser3(ctx context.Context, _headers map[string]string, userId string, photo *string, attrs []int, pattrs *[]int) (_resp *resty.Response, code int, data *string, msg error)
	GetUser4(ctx context.Context, _headers map[string]string, userId string, photo *string, pattrs *[]int, attrs2 ...int) (_resp *resty.Response, code int, data *string, msg error)
	PageUsers3(ctx context.Context, _headers map[string]string, query vo.PageQuery1) (_resp *resty.Response, code int, data vo.PageRet, msg error)
}
