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
		err   error
	)
	ctx = _req.Context()
	if err := json.NewDecoder(_req.Body).Decode(&query); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer _req.Body.Close()
	code, data, err = receiver.ordersvc.PageUsers(
		ctx,
		query,
	)
	if err != nil {
		if err == context.Canceled {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
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

func (receiver *OrdersvcHandlerImpl) GetHello(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		ret string
		err error
	)
	ctx = _req.Context()
	ret, err = receiver.ordersvc.GetHello(
		ctx,
	)
	if err != nil {
		if err == context.Canceled {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(_writer).Encode(struct {
		Ret string `json:"ret,omitempty"`
	}{
		Ret: ret,
	}); err != nil {
		http.Error(_writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (receiver *OrdersvcHandlerImpl) GetHelloWorld(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		ret string
		err error
	)
	ctx = _req.Context()
	ret, err = receiver.ordersvc.GetHelloWorld(
		ctx,
	)
	if err != nil {
		if err == context.Canceled {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(_writer).Encode(struct {
		Ret string `json:"ret,omitempty"`
	}{
		Ret: ret,
	}); err != nil {
		http.Error(_writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (receiver *OrdersvcHandlerImpl) GetGreeting(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx   context.Context
		hello string
		ret   string
		err   error
	)
	ctx = _req.Context()
	hello = _req.FormValue("hello")
	ret, err = receiver.ordersvc.GetGreeting(
		ctx,
		hello,
	)
	if err != nil {
		if err == context.Canceled {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := json.NewEncoder(_writer).Encode(struct {
		Ret string `json:"ret,omitempty"`
	}{
		Ret: ret,
	}); err != nil {
		http.Error(_writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
