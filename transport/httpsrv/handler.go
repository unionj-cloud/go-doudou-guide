package httpsrv

import (
	"net/http"

	ddhttp "github.com/unionj-cloud/go-doudou/svc/http"
)

type UsersvcHandler interface {
	PageUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	UploadAvatar(w http.ResponseWriter, r *http.Request)
	DownloadAvatar(w http.ResponseWriter, r *http.Request)
}

func Routes(handler UsersvcHandler) []ddhttp.Route {
	return []ddhttp.Route{
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
			"GET",
			"/usersvc/downloadavatar",
			handler.DownloadAvatar,
		},
	}
}
