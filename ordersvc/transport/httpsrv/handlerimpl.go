package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"
	service "ordersvc"
	"ordersvc/vo"

	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
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
	if _req.Body == nil {
		http.Error(_writer, "missing request body", http.StatusBadRequest)
		return
	} else {
		if _err := json.NewDecoder(_req.Body).Decode(&query); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	}
	code, data, err = receiver.ordersvc.PageUsers(
		ctx,
		query,
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else if _err, ok := err.(*ddhttp.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Code int        `json:"code"`
		Data vo.PageRet `json:"data"`
	}{
		Code: code,
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
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
		if errors.Is(err, context.Canceled) {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else if _err, ok := err.(*ddhttp.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Ret string `json:"ret"`
	}{
		Ret: ret,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
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
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["hello"]; exists {
		hello = _req.FormValue("hello")
	} else {
		http.Error(_writer, "missing parameter hello", http.StatusBadRequest)
		return
	}
	ret, err = receiver.ordersvc.GetGreeting(
		ctx,
		hello,
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else if _err, ok := err.(*ddhttp.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Ret string `json:"ret"`
	}{
		Ret: ret,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
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
		if errors.Is(err, context.Canceled) {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else if _err, ok := err.(*ddhttp.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Ret string `json:"ret"`
	}{
		Ret: ret,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewOrdersvcHandler(ordersvc service.Ordersvc) OrdersvcHandler {
	return &OrdersvcHandlerImpl{
		ordersvc,
	}
}
