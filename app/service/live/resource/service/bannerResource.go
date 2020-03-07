package service

import (
	"context"
	"github.com/longhao/music/app/service/live/resource/conf"
	"github.com/longhao/music/app/service/live/resource/dao"
	pb "github.com/longhao/music/app/service/live/resource/api/grpc"
)
type BannerResourceService struct {
	dao *dao.Dao
}

func NewBannerResourceService(c *conf.Config) (s *BannerResourceService) {
	s = &BannerResourceService{
		dao: dao.New(c),
	}
	return
}

// List 请求列表资源
func (s *BannerResourceService) List(ctx context.Context, req *pb.BannerListReq) (resp *pb.BannerListResp, err error) {
	var page int32 = 1
	var pageSize int32 = 50

	if req.Page > 0 {
		page = req.Page
	}

	if req.PageSize > 0 {
		pageSize = req.PageSize
	}

	resp = &pb.BannerListResp{}
	resp.CurrentPage = page
	resp.TotalCount, _ = s.dao.GetTotalCount(ctx)

	list, err := s.dao.ListBannerResourceInfo(ctx, page, pageSize)
	if len(list) <= 0 {
		return
	}

	for _, v := range list {
		item := &pb.BannerListResp_List{}
		item.Pic = v.Pic

		resp.List = append(resp.List, item)
	}

	return
}