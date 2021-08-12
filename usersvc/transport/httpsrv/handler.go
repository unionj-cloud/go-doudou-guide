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
	DownloadAvatar(w http.ResponseWriter, r *http.Request)
}

func Routes(handler UsersvcHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"PageUsers",
			"POST",
			"/usersvc/pageusers",
			handler.PageUsers,
		},
		{
			"User",
			"GET",
			"/usersvc/user",
			handler.GetUser,
		},
		{
			"SignUp",
			"POST",
			"/usersvc/signup",
			handler.SignUp,
		},
		{
			"UploadAvatar",
			"POST",
			"/usersvc/uploadavatar",
			handler.UploadAvatar,
		},
		{
			"DownloadAvatar",
			"POST",
			"/usersvc/downloadavatar",
			handler.DownloadAvatar,
		},
	}
}
