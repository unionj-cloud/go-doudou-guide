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

	_cast "github.com/unionj-cloud/cast"
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
		if msg == context.Canceled {
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
	userId = _req.FormValue("userId")
	photo = _req.FormValue("photo")
	code, data, msg = receiver.usersvc.GetUser(
		ctx,
		userId,
		photo,
	)
	if msg != nil {
		if msg == context.Canceled {
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
	username = _req.FormValue("username")
	if casted, err := _cast.ToIntE(_req.FormValue("password")); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	} else {
		password = casted
	}
	if casted, err := _cast.ToBoolE(_req.FormValue("actived")); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	} else {
		actived = casted
	}
	if casted, err := _cast.ToFloat64E(_req.FormValue("score")); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	} else {
		score = casted
	}
	code, data, msg = receiver.usersvc.SignUp(
		ctx,
		username,
		password,
		actived,
		score,
	)
	if msg != nil {
		if msg == context.Canceled {
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
		pf []*multipart.FileHeader
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
	pfFiles := _req.MultipartForm.File["pf"]
	pf = pfFiles
	ps = _req.FormValue("ps")
	ri, rs, re = receiver.usersvc.UploadAvatar(
		pc,
		pf,
		ps,
	)
	if re != nil {
		if re == context.Canceled {
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
func (receiver *UsersvcHandlerImpl) GetDownloadAvatar(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		userId string
		rs     string
		rf     *os.File
		re     error
	)
	ctx = _req.Context()
	userId = _req.FormValue("userId")
	rs, rf, re = receiver.usersvc.GetDownloadAvatar(
		ctx,
		userId,
	)
	if re != nil {
		if re == context.Canceled {
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
	var _fi os.FileInfo
	_fi, _err := rf.Stat()
	if _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
	_writer.Header().Set("Content-Disposition", "inline; filename="+_fi.Name())
	_writer.Header().Set("Content-Type", rs)
	_writer.Header().Set("Content-Length", fmt.Sprintf("%d", _fi.Size()))
	io.Copy(_writer, rf)
}

func NewUsersvcHandler(usersvc service.Usersvc) UsersvcHandler {
	return &UsersvcHandlerImpl{
		usersvc,
	}
}

func (receiver *UsersvcHandlerImpl) UploadAvatar2(_writer http.ResponseWriter, _req *http.Request) {
	var (
		pc  context.Context
		pf  []*multipart.FileHeader
		ps  string
		pf2 *multipart.FileHeader
		pf3 *multipart.FileHeader
		ri  int
		rs  string
		re  error
	)
	pc = _req.Context()
	if err := _req.ParseMultipartForm(32 << 20); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	pfFiles := _req.MultipartForm.File["pf"]
	pf = pfFiles
	ps = _req.FormValue("ps")
	if err := _req.ParseMultipartForm(32 << 20); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	pf2Files := _req.MultipartForm.File["pf2"]
	if len(pf2Files) > 0 {
		pf2 = pf2Files[0]
	}
	if err := _req.ParseMultipartForm(32 << 20); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	pf3Files := _req.MultipartForm.File["pf3"]
	if len(pf3Files) > 0 {
		pf3 = pf3Files[0]
	}
	ri, rs, re = receiver.usersvc.UploadAvatar2(
		pc,
		pf,
		ps,
		pf2,
		pf3,
	)
	if re != nil {
		if re == context.Canceled {
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
