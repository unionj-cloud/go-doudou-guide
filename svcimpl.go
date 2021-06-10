package service

import (
	"context"
	"github.com/go-resty/resty/v2"
	"mime/multipart"
	"os"
	"strings"
	"usersvc/config"
	"usersvc/vo"

	"github.com/jmoiron/sqlx"
)

type UsersvcImpl struct {
	conf config.Config
}

func (receiver *UsersvcImpl) GetUser(ctx context.Context, userId string, photo string) (code int, data string, msg error) {
	panic("implement me")
}

func (receiver *UsersvcImpl) DownloadAvatar(ctx context.Context, userId string) (file *os.File, msg error) {
	downloadLink := "http://upload.wikimedia.org/wikipedia/en/b/bc/Wiki.png"
	splits := strings.Split(downloadLink, "/")
	fileName := splits[len(splits)-1]

	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
	// Setting output directory path, If directory not exists then resty creates one!
	// This is optional one, if you're planning using absoule path in
	// `Request.SetOutput` and can used together.
	client.SetOutputDirectory("tmp")

	// HTTP response gets saved into file, similar to curl -o flag
	_, err := client.R().
		SetOutput(fileName).
		Get(downloadLink)
	if err != nil {
		return nil, err
	}

	return os.Open("tmp/" + fileName)
}

func (receiver *UsersvcImpl) SignUp(ctx context.Context, username string, password string) (code int, data string, msg error) {
	panic("implement me")
}

func (receiver *UsersvcImpl) UploadAvatar(ctx context.Context, avatar []*multipart.FileHeader, userId string) (code int, data string, msg error) {
	panic("implement me")
}

func (receiver *UsersvcImpl) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error) {
	return
}

func NewUsersvc(conf config.Config, db *sqlx.DB) Usersvc {
	return &UsersvcImpl{
		conf,
	}
}
