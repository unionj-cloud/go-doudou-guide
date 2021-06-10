package httpsrv

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	service "usersvc"
	"usersvc/vo"
)

type UsersvcHandlerImpl struct {
	usersvc service.Usersvc
}

func (receiver *UsersvcHandlerImpl) PageUsers(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   context.Context
		query vo.PageQuery
		code  int
		data  vo.PageRet
		msg   error
	)

	ctx = context.Background()

	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	code, data, msg = receiver.usersvc.PageUsers(
		ctx,
		query,
	)

	if msg != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(struct {
		Code int        `json:"code"`
		Data vo.PageRet `json:"data"`
		Msg  error      `json:"msg"`
	}{
		Code: code,
		Data: data,
		Msg:  msg,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}
func (receiver *UsersvcHandlerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    context.Context
		userId string
		photo  string
		code   int
		data   string
		msg    error
	)

	ctx = context.Background()

	userId = r.FormValue("userId")

	photo = r.FormValue("photo")

	code, data, msg = receiver.usersvc.GetUser(
		ctx,
		userId,
		photo,
	)

	if msg != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(struct {
		Code int    `json:"code"`
		Data string `json:"data"`
		Msg  error  `json:"msg"`
	}{
		Code: code,
		Data: data,
		Msg:  msg,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}
func (receiver *UsersvcHandlerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      context.Context
		username string
		password string
		code     int
		data     string
		msg      error
	)

	ctx = context.Background()

	username = r.FormValue("username")

	password = r.FormValue("password")

	code, data, msg = receiver.usersvc.SignUp(
		ctx,
		username,
		password,
	)

	if msg != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(struct {
		Code int    `json:"code"`
		Data string `json:"data"`
		Msg  error  `json:"msg"`
	}{
		Code: code,
		Data: data,
		Msg:  msg,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}
func (receiver *UsersvcHandlerImpl) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	var (
		pc context.Context
		pf []*multipart.FileHeader
		ps string
		ri int
		rs string
		re error
	)

	pc = context.Background()

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	files := r.MultipartForm.File["pf"]
	if len(files) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no file received"))
		return
	}

	pf = files

	ps = r.FormValue("ps")

	ri, rs, re = receiver.usersvc.UploadAvatar(
		pc,
		pf,
		ps,
	)

	if re != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(struct {
		Ri int    `json:"ri"`
		Rs string `json:"rs"`
		Re error  `json:"re"`
	}{
		Ri: ri,
		Rs: rs,
		Re: re,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}
func (receiver *UsersvcHandlerImpl) DownloadAvatar(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    context.Context
		userId string
		rf     *os.File
		re     error
	)

	ctx = context.Background()

	userId = r.FormValue("userId")

	rf, re = receiver.usersvc.DownloadAvatar(
		ctx,
		userId,
	)

	if re != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	fi, err := rf.Stat()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fi.Name())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fi.Size()))

	io.Copy(w, rf)

}

func NewUsersvcHandler(usersvc service.Usersvc) UsersvcHandler {
	return &UsersvcHandlerImpl{
		usersvc,
	}
}
