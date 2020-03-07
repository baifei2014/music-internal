package service

import (
	svcRspb "github.com/longhao/music/app/service/live/resource/api/grpc"
)

type Service struct {
	svcCli *svcRspb.Client
}

func New() *Service {
	s := &Service{
		svcCli: svcRspb.NewClient(),
	}
	return s
}