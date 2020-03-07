package service

import (
	"context"
	"github.com/longhao/music/app/service/main/music/conf"
	"github.com/longhao/music/app/service/main/music/dao"
	pb "github.com/longhao/music/app/service/main/music/api/grpc"
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

func (s *UserResourceService) ListUserInfo(ctx context.Context, req *pb.ListUserInfoReq) (resp *pb.ListUserInfoResp, err error) {
	userInfos, err := s.dao.BatchUserInfo(ctx, req.Uid);

	resp = &pb.ListUserInfoResp{}
	for _, v := range userInfos {
		elem := &pb.ListUserInfoResp_List{
			Uid: v.Uid, 
			Nickname: v.Nickname,
			AvatarUrl: v.AvatarUrl,
		}

		resp.List = append(resp.List, elem)
	}
	return
}