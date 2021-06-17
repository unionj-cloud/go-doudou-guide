package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"
	service "ordersvc"
	"ordersvc/vo"
)

type OrdersvcHandlerImpl struct {
	ordersvc service.Ordersvc
}

func (receiver *OrdersvcHandlerImpl) PageUsers(_writer http.ResponseWriter, _req *http.Request) {
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
	code, data, msg = receiver.ordersvc.PageUsers(
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
		Code int        `json:"code,omitempty"`
		Data vo.PageRet `json:"data,omitempty"`
	}{
		Code: code,
		Data: data,
	}); err != nil {
		http.Error(_writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewOrdersvcHandler(ordersvc service.Ordersvc) OrdersvcHandler {
	return &OrdersvcHandlerImpl{
		ordersvc,
	}
}
