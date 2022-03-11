package service

import (
	"context"
	"statsvc/config"
	"statsvc/internal/reportsvcj"
)

type StatsvcImpl struct {
	conf       *config.Config
	echoClient *reportsvcj.EchoClient
}

func (receiver *StatsvcImpl) Add(ctx context.Context, x int, y int) (data int, err error) {
	return x + y, nil
}

func NewStatsvc(conf *config.Config, echoClient *reportsvcj.EchoClient) Statsvc {
	return &StatsvcImpl{
		conf,
		echoClient,
	}
}

func (receiver *StatsvcImpl) GetEcho(ctx context.Context, s string) (data string, err error) {
	data, _, err = receiver.echoClient.GetEchoString(ctx, nil, s)
	return
}
