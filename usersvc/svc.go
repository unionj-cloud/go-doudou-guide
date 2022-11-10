package service

import (
	"context"
	v3 "github.com/unionj-cloud/go-doudou/v2/toolkit/openapi/v3"
	"os"
	"usersvc/vo"
)

//go:generate go-doudou svc http -c
//go:generate go-doudou svc grpc

// Usersvc User Center Service
type Usersvc interface {
	// PageUsers demonstrate how to define POST and Content-Type as application/json api
	PageUsers(ctx context.Context,
		// comments above input and output struct type parameters in vo package will display on online document
		// not comments here
		query vo.PageQuery) (
		// it indicates there is an error if code is not 0
		code int,
		// page data
		data vo.PageRet,
		// error message
		msg error)

	// GetUser demonstrate how to define GET api with query string parameters
	GetUser(ctx context.Context,
		// user id
		// comments above input and output basic type parameters will display on online document
		userId string, photo string) (code int, data string, msg error)

	// SignUp demonstrate how to define POST and Content-Type as application/x-www-form-urlencoded api
	SignUp(ctx context.Context, username string, password int, actived bool, score float64) (code int, data string, msg error)

	// UploadAvatar demonstrate how to define upload files api
	// there must be one []v3.FileModel or v3.FileModel parameter among input parameters
	// remember to close the readers by Close method of v3.FileModel if you don't need them anymore when you finished your own business logic
	UploadAvatar(context.Context, []v3.FileModel, string) (int, string, error)

	// UploadAvatar2 demonstrate how to define upload files api
	// remember to close the readers by Close method of v3.FileModel if you don't need them anymore when you finished your own business logic
	UploadAvatar2(context.Context, []v3.FileModel, string, *v3.FileModel, *v3.FileModel) (int, string, error)

	// GetDownloadAvatar demonstrate how to define download file api
	// there must be *os.File parameter among output parameters
	GetDownloadAvatar(ctx context.Context, userId string) (
		// mimetype(Content-Type) for download file
		// you don't have to add this parameter because you will have a default value application/octet-stream
		// in generated handlerimpl.go file.
		// if you add it, you should make a small fix manually to replace application/octet-stream with it.
		// Any custom manual fix in handlerimpl.go file won't been overwritten when you re-execute go-doudou commands
		// like go-doudou svc http --handler -c go --doc
		// go-doudou will ignore any output parameter other than *os.File when generate OpenAPI 3.0 json file and online document
		string,
		// download file
		*os.File, error)

	// GetUser2 demonstrate how to define GET api with query string parameters
	GetUser2(ctx context.Context,
		// user id
		// comments above input and output basic type parameters will display on online document
		userId string, photo *string) (code int, data *string, msg error)

	// PageUsers2 demonstrate how to define POST and Content-Type as application/json api
	PageUsers2(ctx context.Context,
		// comments above input and output struct type parameters in vo package will display on online document
		// not comments here
		query *vo.PageQuery) (
		// it indicates there is an error if code is not 0
		code int,
		// page data
		data vo.PageRet,
		// error message
		msg error)

	// GetUser3 demonstrate how to define GET api with query string parameters
	GetUser3(ctx context.Context,
		// user id
		// comments above input and output basic type parameters will display on online document
		userId string, photo *string, attrs []int, pattrs *[]int) (code int, data *string, msg error)

	// GetUser4 demonstrate how to define GET api with query string parameters
	// photo *string, pattrs *[]int 是一类问题， TODO
	// attrs2 []int 是一类问题
	GetUser4(ctx context.Context,
		// user id
		// comments above input and output basic type parameters will display on online document
		userId string, photo *string, pattrs *[]int, attrs2 ...int) (code int, data *string, msg error)

	// PageUsers3 demonstrate how to define POST and Content-Type as application/json api
	PageUsers3(ctx context.Context,
		// comments above input and output struct type parameters in vo package will display on online document
		// not comments here
		query vo.PageQuery1) (
		// it indicates there is an error if code is not 0
		code int,
		// page data
		data vo.PageRet,
		// error message
		msg error)
}
