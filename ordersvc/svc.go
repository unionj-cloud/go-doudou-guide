package service

import (
	"context"
	"ordersvc/vo"
)

type Ordersvc interface {
	// You can define your service methods as your need. Below is an example.
	PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, err error)
	GetHello(ctx context.Context) (ret string, err error)
	GetGreeting(ctx context.Context, hello string) (ret string, err error)
	GetHelloWorld(ctx context.Context) (ret string, err error)
}
