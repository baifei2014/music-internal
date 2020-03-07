package service

import (
	"context"

	"github.com/longhao/music/app/interface/main/user/model"

	rspb "github.com/longhao/music/app/service/live/resource/api/grpc"
)

func (s *Service) GetUserSingle(ctx context.Context, uid int32) (resp *model.UserResourceGetSingleResp, err error) {
	respRPC, err := s.svcCli.Query(ctx, &rspb.QueryReq{
		Uid: uid,
	})

	if err == nil {
		resp = &model.UserResourceGetSingleResp{
			Uid: respRPC.Uid,
			Username: respRPC.Username,
		}
	}
	return
}