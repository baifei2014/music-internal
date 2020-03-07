package music

import (
	"fmt"
	"context"

	"github.com/longhao/music/app/interface/main/user/model"
	"github.com/longhao/music/app/interface/main/user/model/music"

	"github.com/longhao/music/library/xstr"

	rspb "github.com/longhao/music/app/service/main/music/api/grpc"
)

func (s *Service) UserPlaylistGet(ctx context.Context, userId int64, page int64, pageSize int64) (resp *music.UserPlaylistResp, err error) {
	resp = &music.UserPlaylistResp{
		More: false,
		Code: 200,
		Playlist: []*music.Playlist{},
	}

	respRPC, err := s.svcCli.UserPlaylistGet(ctx, &rspb.UserPlaylistGetReq{
		UserId: userId,
		Page: page,
		PageSize: pageSize,
	})
	fmt.Println(respRPC)

	if len(respRPC.Playlist) <= 0 {
		return
	}

	if respRPC.TotalCount > respRPC.CurrentPage * pageSize {
		resp.More = true
	}

	var userIds []int64
	userIds = getUserIdInPlaylist(respRPC.Playlist)

	userInfos, _ := s.batchUserInfo(ctx, userIds)

	for _, v := range respRPC.Playlist {
		elem := &music.Playlist{
			Id: v.Id,
			Name: v.Name,
			CoverImgUrl: xstr.GetFullImageUrl(v.CoverImgUrl),
			Creator: userInfos[v.CreatorUid],
			PlayCount: v.PlayCount,
			CommentCount: v.CommentCount,
			ShareCount: v.ShareCount,
			TrackCount: v.TrackCount,
		}
		resp.Playlist = append(resp.Playlist, elem)
	}

	return
}

func (s *Service) batchUserInfo(c context.Context, uids []int64) (resp map[int64]*model.UserBase, err error) {
	resp, _ = s.dao.BatchUserInfo(c, uids)
	return
}

func getUserIdInPlaylist(p []*rspb.UserPlaylistGetResp_Playlist) (l []int64) {
	for _, v := range p {
		l = append(l, v.CreatorUid)
	}

	return
}