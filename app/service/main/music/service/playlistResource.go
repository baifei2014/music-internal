package service

import (
	"fmt"
	"context"

	"github.com/longhao/music/app/service/main/music/dao"
	"github.com/longhao/music/app/service/main/music/conf"
	pb "github.com/longhao/music/app/service/main/music/api/grpc"
)

type PlaylistResourceService struct {
	dao *dao.Dao
}

func NewPlaylistResourceService(c *conf.Config) (s *PlaylistResourceService) {
	s = &PlaylistResourceService{
		dao: dao.New(c),
	}
	return
}

func (s *PlaylistResourceService) PlaylistDetailGet(ctx context.Context, req *pb.PlaylistDetailGetReq) (resp *pb.PlaylistDetailGetResp, err error) {
	playlistInfo, err := s.dao.GetPlaylistDetailInfo(ctx, req.Id)

	fmt.Println("service查询结果")
	fmt.Println(playlistInfo)
	resp = &pb.PlaylistDetailGetResp{
		Id: playlistInfo.Id,
		Name: playlistInfo.Name,
		CoverImgUrl: playlistInfo.CoverImgUrl,
		CreatorUid: playlistInfo.CreatorUid,
		PlayCount: playlistInfo.PlayCount,
		CommentCount: playlistInfo.CommentCount,
		ShareCount: playlistInfo.ShareCount,
		TrackCount: playlistInfo.TrackCount,
		SubscribedCount: playlistInfo.SubscribedCount,
		Sids: playlistInfo.Sids,
	}

	fmt.Println(resp)

	return
}

func (s *PlaylistResourceService) UserPlaylistGet(ctx context.Context, req *pb.UserPlaylistGetReq) (resp *pb.UserPlaylistGetResp, err error) {
	var (
		Page int64 = 1
		PageSize int64 = 1000
	)
	if req.Page > 0 {
		Page = req.Page
	}
	if req.PageSize > 0 {
		PageSize = req.PageSize
	}
	resp = &pb.UserPlaylistGetResp{}

	count, _ := s.dao.GetTotalCount(ctx, req.UserId)
	resp.CurrentPage = Page
	resp.TotalCount = count
	list, err := s.dao.ListUserPlaylistResourceInfo(ctx, req.UserId, Page, PageSize)

	fmt.Println(resp)
	if len(list) <= 0 {
		return
	}

	for _, v := range list {
		elem := &pb.UserPlaylistGetResp_Playlist{
			Id: v.Id,
			Name: v.Name,
			CoverImgUrl: v.CoverImgUrl,
			CreatorUid: v.CreatorUid,
			PlayCount: v.PlayCount,
			CommentCount: v.CommentCount,
			ShareCount: v.ShareCount,
			TrackCount: v.TrackCount,
		}
		resp.Playlist = append(resp.Playlist, elem)
	}

	fmt.Println(resp)
	return
}