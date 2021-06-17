package service

import (
	"admin/config"
	"admin/vo"
	"context"

	"github.com/jmoiron/sqlx"
)

type AdminImpl struct {
	conf config.Config
}

func (receiver *AdminImpl) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, msg error) {
	panic("implement me")
}

func NewAdmin(conf config.Config, db *sqlx.DB) Admin {
	return &AdminImpl{
		conf,
	}
}
