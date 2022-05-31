package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"ordersvc/config"
	"ordersvc/vo"
	userclient "usersvc/client"
)

type OrdersvcImpl struct {
	conf          *config.Config
	usersvcClient userclient.IUsersvcClient
}

func (receiver *OrdersvcImpl) GetGreeting(ctx context.Context, hello string) (ret string, err error) {
	return hello, nil
}

func (receiver *OrdersvcImpl) GetHelloWorld(ctx context.Context) (ret string, err error) {
	return "Hello World", nil
}

func (receiver *OrdersvcImpl) GetHello(ctx context.Context) (ret string, err error) {
	return "world", nil
}

func (receiver *OrdersvcImpl) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error) {
	//var _data vo1.PageRet
	//_, code, _data, msg = receiver.usersvcClient.PageUsers(ctx, nil, vo1.PageQuery{
	//	Filter: vo1.PageFilter(query.Filter),
	//	Page: vo1.Page{
	//		Orders: nil,
	//		PageNo: query.Page.PageNo,
	//		Size:   query.Page.Size,
	//	},
	//})
	//data = vo.PageRet(_data)
	return
}

func NewOrdersvc(conf *config.Config, db *sqlx.DB, usersvcClient userclient.IUsersvcClient) Ordersvc {
	return &OrdersvcImpl{
		conf,
		usersvcClient,
	}
}
