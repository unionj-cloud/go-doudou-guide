package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type UsersvcHandler interface {
	PageUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	UploadAvatar(w http.ResponseWriter, r *http.Request)
	UploadAvatar2(w http.ResponseWriter, r *http.Request)
	GetDownloadAvatar(w http.ResponseWriter, r *http.Request)
	GetUser2(w http.ResponseWriter, r *http.Request)
	PageUsers2(w http.ResponseWriter, r *http.Request)
	GetUser3(w http.ResponseWriter, r *http.Request)
	GetUser4(w http.ResponseWriter, r *http.Request)
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
		{
			"User2",
			"GET",
			"/user/2",
			handler.GetUser2,
		},
		{
			"PageUsers2",
			"POST",
			"/page/users/2",
			handler.PageUsers2,
		},
		{
			"User3",
			"GET",
			"/user/3",
			handler.GetUser3,
		},
		{
			"User4",
			"GET",
			"/user/4",
			handler.GetUser4,
		},
	}
}
