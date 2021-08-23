package service

import (
	"context"
	"ordersvc/config"
	"ordersvc/vo"
	service "usersvc"
	vo1 "usersvc/vo"

	"github.com/jmoiron/sqlx"
)

type OrdersvcImpl struct {
	conf          *config.Config
	usersvcClient service.Usersvc
}

func (receiver *OrdersvcImpl) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error) {
	var _data vo1.PageRet
	code, _data, msg = receiver.usersvcClient.PageUsers(ctx, vo1.PageQuery{
		Filter: vo1.PageFilter(query.Filter),
		Page: vo1.Page{
			Orders: nil,
			PageNo: query.Page.PageNo,
			Size:   query.Page.Size,
		},
	})
	data = vo.PageRet(_data)
	return
}

func NewOrdersvc(conf *config.Config, db *sqlx.DB, usersvcClient service.Usersvc) Ordersvc {
	return &OrdersvcImpl{
		conf,
		usersvcClient,
	}
}
