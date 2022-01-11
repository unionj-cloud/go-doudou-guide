package httpsrv

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	service "usersvc"
	"usersvc/vo"

	"github.com/pkg/errors"
	"github.com/unionj-cloud/cast"
	v3 "github.com/unionj-cloud/go-doudou/openapi/v3"
)

type UsersvcHandlerImpl struct {
	usersvc service.Usersvc
}

func (receiver *UsersvcHandlerImpl) PageUsers(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx   context.Context
		query vo.PageQuery
		code  int
		data  vo.PageRet
		msg   error
	)
	ctx = _req.Context()
	if err := json.NewDecoder(_req.Body).Decode(&query); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer _req.Body.Close()
	code, data, msg = receiver.usersvc.PageUsers(
		ctx,
		query,
	)
	if msg != nil {
		if errors.Is(msg, context.Canceled) {
			http.Error(_writer, msg.Error(), http.StatusBadRequest)
		} else {
			http.Error(_writer, msg.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(_writer).Encode(struct {
		Code int        `json:"code"`
		Data vo.PageRet `json:"data"`
	}{
		Code: code,
		Data: data,
	}); err != nil {
		http.Error(_writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *UsersvcHandlerImpl) GetUser(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		userId string
		photo  string
		code   int
		data   string
		msg    error
	)
	ctx = _req.Context()
	if err := _req.ParseForm(); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["userId"]; exists {
		userId = _req.FormValue("userId")
	}
	if _, exists := _req.Form["photo"]; exists {
		photo = _req.FormValue("photo")
	}
	code, data, msg = receiver.usersvc.GetUser(
		ctx,
		userId,
		photo,
	)
	if msg != nil {
		if errors.Is(msg, context.Canceled) {
			http.Error(_writer, msg.Error(), http.StatusBadRequest)
		} else {
			http.Error(_writer, msg.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(_writer).Encode(struct {
		Code int    `json:"code"`
		Data string `json:"data"`
	}{
		Code: code,
		Data: data,
	}); err != nil {
		http.Error(_writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *UsersvcHandlerImpl) SignUp(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx      context.Context
		username string
		password int
		actived  bool
		score    float64
		code     int
		data     string
		msg      error
	)
	ctx = _req.Context()
	if err := _req.ParseForm(); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["username"]; exists {
		username = _req.FormValue("username")
	}
	if _, exists := _req.Form["password"]; exists {
		if casted, err := cast.ToIntE(_req.FormValue("password")); err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		} else {
			password = casted
		}
	}
	if _, exists := _req.Form["actived"]; exists {
		if casted, err := cast.ToBoolE(_req.FormValue("actived")); err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		} else {
			actived = casted
		}
	}
	if _, exists := _req.Form["score"]; exists {
		if casted, err := cast.ToFloat64E(_req.FormValue("score")); err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		} else {
			score = casted
		}
	}
	code, data, msg = receiver.usersvc.SignUp(
		ctx,
		username,
		password,
		actived,
		score,
	)
	if msg != nil {
		if errors.Is(msg, context.Canceled) {
			http.Error(_writer, msg.Error(), http.StatusBadRequest)
		} else {
			http.Error(_writer, msg.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(_writer).Encode(struct {
		Code int    `json:"code"`
		Data string `json:"data"`
	}{
		Code: code,
		Data: data,
	}); err != nil {
		http.Error(_writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *UsersvcHandlerImpl) UploadAvatar(_writer http.ResponseWriter, _req *http.Request) {
	var (
		pc context.Context
		pf []*v3.FileModel
		ps string
		ri int
		rs string
		re error
	)
	pc = _req.Context()
	if err := _req.ParseMultipartForm(32 << 20); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	pfFileHeaders := _req.MultipartForm.File["pf"]
	for _, _fh := range pfFileHeaders {
		_f, err := _fh.Open()
		if err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		}
		pf = append(pf, &v3.FileModel{
			Filename: _fh.Filename,
			Reader:   _f,
		})
	}
	if err := _req.ParseForm(); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["ps"]; exists {
		ps = _req.FormValue("ps")
	}
	ri, rs, re = receiver.usersvc.UploadAvatar(
		pc,
		pf,
		ps,
	)
	if re != nil {
		if errors.Is(re, context.Canceled) {
			http.Error(_writer, re.Error(), http.StatusBadRequest)
		} else {
			http.Error(_writer, re.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(_writer).Encode(struct {
		Ri int    `json:"ri"`
		Rs string `json:"rs"`
	}{
		Ri: ri,
		Rs: rs,
	}); err != nil {
		http.Error(_writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *UsersvcHandlerImpl) UploadAvatar2(_writer http.ResponseWriter, _req *http.Request) {
	var (
		pc  context.Context
		pf  []*v3.FileModel
		ps  string
		pf2 *v3.FileModel
		pf3 *v3.FileModel
		ri  int
		rs  string
		re  error
	)
	pc = _req.Context()
	if err := _req.ParseMultipartForm(32 << 20); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	pfFileHeaders := _req.MultipartForm.File["pf"]
	for _, _fh := range pfFileHeaders {
		_f, err := _fh.Open()
		if err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		}
		pf = append(pf, &v3.FileModel{
			Filename: _fh.Filename,
			Reader:   _f,
		})
	}
	if err := _req.ParseForm(); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["ps"]; exists {
		ps = _req.FormValue("ps")
	}
	pf2FileHeaders := _req.MultipartForm.File["pf2"]
	if len(pf2FileHeaders) > 0 {
		_fh := pf2FileHeaders[0]
		_f, err := _fh.Open()
		if err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		}
		pf2 = &v3.FileModel{
			Filename: _fh.Filename,
			Reader:   _f,
		}
	}
	pf3FileHeaders := _req.MultipartForm.File["pf3"]
	if len(pf3FileHeaders) > 0 {
		_fh := pf3FileHeaders[0]
		_f, err := _fh.Open()
		if err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		}
		pf3 = &v3.FileModel{
			Filename: _fh.Filename,
			Reader:   _f,
		}
	}
	ri, rs, re = receiver.usersvc.UploadAvatar2(
		pc,
		pf,
		ps,
		pf2,
		pf3,
	)
	if re != nil {
		if errors.Is(re, context.Canceled) {
			http.Error(_writer, re.Error(), http.StatusBadRequest)
		} else {
			http.Error(_writer, re.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(_writer).Encode(struct {
		Ri int    `json:"ri"`
		Rs string `json:"rs"`
	}{
		Ri: ri,
		Rs: rs,
	}); err != nil {
		http.Error(_writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewUsersvcHandler(usersvc service.Usersvc) UsersvcHandler {
	return &UsersvcHandlerImpl{
		usersvc,
	}
}

func (receiver *UsersvcHandlerImpl) GetDownloadAvatar(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		userId string
		rf     *os.File
		re     error
	)
	ctx = _req.Context()
	if err := _req.ParseForm(); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["userId"]; exists {
		userId = _req.FormValue("userId")
	}
	_, rf, re = receiver.usersvc.GetDownloadAvatar(
		ctx,
		userId,
	)
	if re != nil {
		if errors.Is(re, context.Canceled) {
			http.Error(_writer, re.Error(), http.StatusBadRequest)
		} else {
			http.Error(_writer, re.Error(), http.StatusInternalServerError)
		}
		return
	}
	if rf == nil {
		http.Error(_writer, "No file returned", http.StatusInternalServerError)
		return
	}
	defer rf.Close()
	var _fi os.FileInfo
	_fi, _err := rf.Stat()
	if _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
	_writer.Header().Set("Content-Disposition", "attachment; filename="+_fi.Name())
	_writer.Header().Set("Content-Type", "application/octet-stream")
	_writer.Header().Set("Content-Length", fmt.Sprintf("%d", _fi.Size()))
	io.Copy(_writer, rf)
}
