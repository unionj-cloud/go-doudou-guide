package service

import (
	"context"
	"mime/multipart"
	"os"
	"usersvc/vo"
)

// Usersvc 用户服务接口
// v1版本
type Usersvc interface {
	// PageUsers You can define your service methods as your need. Below is an example.
	PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error)

	// GetUser comment1
	// comment2
	GetUser(ctx context.Context,
	// 用户ID
		userId string,
	// 图片地址
		photo string,
	) (code int, data string, msg error)

	// SignUp comment3
	SignUp(ctx context.Context, username string, password int, actived bool, score float64) (code int, data string, msg error)

	// UploadAvatar comment4
	UploadAvatar(context.Context, []*multipart.FileHeader, string) (int, string, error)

	// DownloadAvatar comment5
	DownloadAvatar(ctx context.Context, userId string) (*os.File, error)
}
