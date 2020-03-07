package chat

import (
	"github.com/longhao/music/app/interface/main/user/dao"
	svcRspb "github.com/longhao/music/app/service/main/music/api/grpc"
)

type Service struct {
	svcCli *svcRspb.Client
	dao *dao.Dao
}

func New() *Service {
	s := &Service{
		svcCli: svcRspb.NewClient(),
		dao: dao.New(),
	}
	return s
}