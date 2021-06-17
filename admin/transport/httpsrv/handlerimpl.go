package httpsrv

import (
	service "admin"
	"net/http"
)

type AdminHandlerImpl struct {
	admin service.Admin
}

func (receiver *AdminHandlerImpl) PageUsers(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func NewAdminHandler(admin service.Admin) AdminHandler {
	return &AdminHandlerImpl{
		admin,
	}
}
