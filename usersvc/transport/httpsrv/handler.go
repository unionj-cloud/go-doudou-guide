package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/svc/http/model"
)

type UsersvcHandler interface {
	PageUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	UploadAvatar(w http.ResponseWriter, r *http.Request)
	UploadAvatar2(w http.ResponseWriter, r *http.Request)
	GetDownloadAvatar(w http.ResponseWriter, r *http.Request)
}

func Routes(handler UsersvcHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"PageUsers",
			"POST",
			"/page/users",
			handler.PageUsers,
		},
		{
			"User",
			"GET",
			"/user",
			handler.GetUser,
		},
		{
			"SignUp",
			"POST",
			"/sign/up",
			handler.SignUp,
		},
		{
			"UploadAvatar",
			"POST",
			"/upload/avatar",
			handler.UploadAvatar,
		},
		{
			"UploadAvatar2",
			"POST",
			"/upload/avatar/2",
			handler.UploadAvatar2,
		},
		{
			"DownloadAvatar",
			"GET",
			"/download/avatar",
			handler.GetDownloadAvatar,
		},
	}
}
