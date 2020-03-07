package service

import (
	"context"

	"github.com/longhao/music/app/interface/main/user/model"

	rspb "github.com/longhao/music/app/service/live/resource/api/grpc"
)

func (s *Service) GetBanner(ctx context.Context, page int32, pageSize int32) (resp *model.BannerResourceGetResp, err error) {
	respRPC, err := s.svcCli.List(ctx, &rspb.BannerListReq{
		Page: page,
		PageSize: pageSize,
	})

	if err == nil {
		resp = &model.BannerResourceGetResp{
			CurrentPage: respRPC.CurrentPage,
			TotalCount: respRPC.TotalCount,
			List: convertRPCListRes(respRPC.List),
		}
	}
	return
}

func convertRPCListRes(RPCList []*rspb.BannerListResp_List) (HTTPList []*model.BannerResourceGetResp_List) {
	HTTPList = make([]*model.BannerResourceGetResp_List, len(RPCList))
	for index, RPCListItem := range RPCList {
		HTTPList[index] = &model.BannerResourceGetResp_List{
			Pic: RPCListItem.Pic,
		}
	}
	return
}