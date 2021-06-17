package service

import (
	"context"
	"ordersvc/vo"
)

type Ordersvc interface {
	// You can define your service methods as your need. Below is an example.
	PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error)
}
