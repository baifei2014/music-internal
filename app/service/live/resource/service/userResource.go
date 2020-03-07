package service

import (
	"context"
	"github.com/longhao/music/app/service/live/resource/conf"
	"github.com/longhao/music/app/service/live/resource/dao"
	pb "github.com/longhao/music/app/service/live/resource/api/grpc"
)
type UserResourceService struct {
	dao *dao.Dao
}

func NewUserResourceService(c *conf.Config) (s *UserResourceService) {
	s = &UserResourceService{
		dao: dao.New(c),
	}
	return
}

// Query请求单个资源
func (s *UserResourceService) Query(ctx context.Context, req *pb.QueryReq) (resp *pb.QueryResp, err error) {
	resp = &pb.QueryResp{}

	info, err := s.dao.GetUserResourceInfo(ctx, req.Uid)
	if err != nil {
		return
	}

	resp.Uid = info.Uid
	resp.Username = info.Username
	
	return
}