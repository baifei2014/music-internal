package dao

import (
	rspb "github.com/longhao/music/app/service/main/music/api/grpc"
)

type Dao struct {
	userClient *rspb.Client
}

func New() (dao *Dao) {
	dao = &Dao{
		userClient: rspb.NewClient(),
	}
	return
}