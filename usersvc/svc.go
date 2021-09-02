package service

import (
	"context"
	"mime/multipart"
	"os"
	"usersvc/vo"
)

// Usersvc User Center Service
type Usersvc interface {
	// PageUsers demonstrate how to define POST and Content-Type as application/json api
	PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error)

	// GetUser demonstrate how to define GET api with query string parameters
	GetUser(ctx context.Context, userId string, photo string) (code int, data string, msg error)

	// SignUp demonstrate how to define POST and Content-Type as application/x-www-form-urlencoded api
	SignUp(ctx context.Context, username string, password int, actived bool, score float64) (code int, data string, msg error)

	// UploadAvatar demonstrate how to define upload files api
	// there must be one []*multipart.FileHeader or *multipart.FileHeader parameter among output parameters
	UploadAvatar(context.Context, []*multipart.FileHeader, string) (int, string, error)

	// UploadAvatar demonstrate how to define upload files api
	// there must be one []*multipart.FileHeader or *multipart.FileHeader parameter among output parameters
	UploadAvatar2(context.Context, []*multipart.FileHeader, string, *multipart.FileHeader, *multipart.FileHeader) (int, string, error)

	// GetDownloadAvatar demonstrate how to define download file api
	// there must be *os.File parameter among output parameters
	GetDownloadAvatar(ctx context.Context, userId string) (string, *os.File, error)
}
