package service

import (
	"context"
	"io"
	"os"
	"strings"
	"time"
	"usersvc/config"
	"usersvc/vo"

	v3 "github.com/unionj-cloud/go-doudou/openapi/v3"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	ddconfig "github.com/unionj-cloud/go-doudou/svc/config"
)

type UsersvcImpl struct {
	conf *config.Config
}

func saveFile(fm *v3.FileModel) error {
	defer fm.Close()
	f, err := os.OpenFile(ddconfig.GddOutput.Load()+"/"+fm.Filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "call os.OpenFile error")
	}
	defer f.Close()
	_, err = io.Copy(f, fm.Reader)
	if err != nil {
		return err
	}
	return nil
}

func (receiver *UsersvcImpl) UploadAvatar2(ctx context.Context, headers []*v3.FileModel, s string, header *v3.FileModel, header2 *v3.FileModel) (int, string, error) {
	_ = os.MkdirAll(ddconfig.GddOutput.Load(), os.ModePerm)
	for _, fh := range headers {
		if err := saveFile(fh); err != nil {
			return 1, "", errors.Wrapf(err, "call saveFile error")
		}
	}
	if header != nil {
		if err := saveFile(header); err != nil {
			return 1, "", errors.Wrapf(err, "call saveFile error")
		}
	}
	if header2 != nil {
		if err := saveFile(header2); err != nil {
			return 1, "", errors.Wrapf(err, "call saveFile error")
		}
	}
	return 0, "OK", nil
}

func (receiver *UsersvcImpl) GetDownloadAvatar(ctx context.Context, userId string) (string, *os.File, error) {
	downloadLink := "http://upload.wikimedia.org/wikipedia/en/b/bc/Wiki.png"
	splits := strings.Split(downloadLink, "/")
	fileName := splits[len(splits)-1]

	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
	// Setting output directory path, If directory not exists then resty creates one!
	// This is optional one, if you're planning using absoule path in
	// `Request.SetOutput` and can used together.
	client.SetOutputDirectory(ddconfig.GddOutput.Load())

	// HTTP response gets saved into file, similar to curl -o flag
	resp, err := client.R().
		SetOutput(fileName).
		Get(downloadLink)
	if err != nil {
		return "", nil, err
	}
	mimetype := resp.Header().Get("Content-Type")
	f, err := os.Open(ddconfig.GddOutput.Load() + "/" + fileName)
	return mimetype, f, err
}

func (receiver *UsersvcImpl) UploadAvatar(ctx context.Context, avatar []*v3.FileModel, userId string) (code int, data string, msg error) {
	if len(avatar) == 0 {
		return 1, "", errors.New("no file upload")
	}
	_ = os.MkdirAll(ddconfig.GddOutput.Load(), os.ModePerm)
	err := saveFile(avatar[0])
	if err != nil {
		return 1, "", errors.Wrap(err, "save file failed")
	}
	return 0, "OK", nil
}

func (receiver *UsersvcImpl) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error) {
	select {
	case <-time.After(10 * time.Millisecond):
		//minute := time.Now().Second()
		//if minute%2 != 0 {
		//	panic(fmt.Errorf("error because %d minute is odd", minute))
		//}
		return 0, vo.PageRet{
			Items: []map[string]interface{}{
				{
					"name": "jack",
					"age":  30,
				},
				{
					"name": "rose",
					"age":  23,
				},
			},
			PageNo:   2,
			PageSize: 10,
			Total:    50,
			HasNext:  true,
		}, nil
	case <-ctx.Done():
		msg = ctx.Err()
		code = 1
		return
	}
}

func NewUsersvc(conf *config.Config) Usersvc {
	return &UsersvcImpl{
		conf,
	}
}

func (receiver *UsersvcImpl) GetUser(ctx context.Context, userId string, photo string) (code int, data string, msg error) {
	var _result struct {
		Code int
		Data string
	}
	_ = gofakeit.Struct(&_result)
	return _result.Code, _result.Data, nil
}
func (receiver *UsersvcImpl) SignUp(ctx context.Context, username string, password int, actived bool, score float64) (code int, data string, msg error) {
	var _result struct {
		Code int
		Data string
	}
	_ = gofakeit.Struct(&_result)
	return _result.Code, _result.Data, nil
}
