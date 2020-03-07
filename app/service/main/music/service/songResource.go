package service

import (
	"fmt"
	"context"
	"github.com/longhao/music/app/service/main/music/conf"
	"github.com/longhao/music/app/service/main/music/dao"
	pb "github.com/longhao/music/app/service/main/music/api/grpc"
)
type SongResourceService struct {
	dao *dao.Dao
}

func NewSongResourceService(c *conf.Config) (s *SongResourceService) {
	s = &SongResourceService{
		dao: dao.New(c),
	}
	return
}

func (s *SongResourceService) PlaylistGetList(ctx context.Context, req *pb.PlaylistGetListReq) (resp *pb.PlaylistGetListResp, err error) {
	info, err := s.dao.GetPlaylist(ctx)
	resp = &pb.PlaylistGetListResp {}

	for _, item := range info {
		playlist := &pb.PlaylistGetListResp_Playlist{
			Id: item.Id,
			Name: item.Name,
			CoverImgUrl: item.CoverImgUrl,
		}
		resp.Playlists = append(resp.Playlists, playlist)
	}
	return
}

func (s *SongResourceService) PlaylistGetSingle(ctx context.Context, req *pb.GetPlaylistSingleReq) (resp *pb.GetPlaylistSingleResp, err error) {
	info, err := s.dao.GetPlaylistDetailInfo(ctx, req.Id)

	resp = &pb.GetPlaylistSingleResp{
		Id: info.Id,
		Name: info.Name,
		CoverImgUrl: info.CoverImgUrl,
		PlayCount: info.PlayCount,
		CommentCount: info.CommentCount,
		ShareCount: info.ShareCount,
		Sids: info.Sids,
	}

	return
}

func (s *SongResourceService) ArtistGet(ctx context.Context, req *pb.GetArtistReq) (resp *pb.GetArtistResp, err error) {
	info, err := s.dao.GetArtistInfo(ctx, req.Sids)

	resp = &pb.GetArtistResp{}
	for _, item := range info {
		artist := &pb.GetArtistResp_Artist{
			Id: item.Id,
			Sid: item.Sid,
			Name: item.Name,
		}
		resp.Artists = append(resp.Artists, artist)
	}

	fmt.Println(resp)

	return
}

func (s *SongResourceService) AlbumGet(ctx context.Context, req *pb.GetAlbumReq) (resp *pb.GetAlbumResp, err error) {
	info, err := s.dao.GetAlbumInfo(ctx, req.Sids)

	resp = &pb.GetAlbumResp{}
	for _, albumItem := range info {
		album := &pb.GetAlbumResp_Album{
			Id: albumItem.Id,
			Sid: albumItem.Sid,
			Name: albumItem.Name,
			PicUrl: albumItem.PicUrl,
		}

		resp.Albums = append(resp.Albums, album)
	}

	fmt.Println(resp)

	return
}

// Query请求单个资源
func (s *SongResourceService) SongGetDetail(ctx context.Context, req *pb.GetDetailReq) (resp *pb.GetDetailResp, err error) {
	resp = &pb.GetDetailResp{}

	info, err := s.dao.GetSongResourceInfo(ctx, req.Sids)
	if err != nil {
		return
	}

	privileges, err := s.dao.GetPrivilegeInfo(ctx, req.Sids)
	if err != nil {
		return
	}

	for _, item := range info {
		song := &pb.GetDetailResp_Song{
			Id: item.Sid,
			Name: item.Name,
			Duration: item.Duration,
			Url: item.Url,
		}

		resp.Songs = append(resp.Songs, song)
	}

	for _, item := range privileges {
		privilege := &pb.GetDetailResp_Privilege{
			Id: item.Sid,
			Maxbr: item.Maxbr,
			St: item.St,
		}

		resp.Privileges = append(resp.Privileges, privilege)
	}
	
	fmt.Println(resp)
	
	return
}

// 获取歌曲列表
func (s *SongResourceService) ListSongGet(ctx context.Context, req *pb.ListSongGetReq) (resp *pb.ListSongGetResp, err error) {
	resp = &pb.ListSongGetResp{}

	info, err := s.dao.GetSongResourceInfo(ctx, req.Sids)
	if err != nil {
		return
	}

	for _, item := range info {
		song := &pb.ListSongGetResp_Song{
			Id: item.Sid,
			Name: item.Name,
			Duration: item.Duration,
			Url: item.Url,
		}

		resp.Songs = append(resp.Songs, song)
	}

	privileges, err := s.dao.GetPrivilegeInfo(ctx, req.Sids)
	if err == nil {
		for _, item := range privileges {
			privilege := &pb.ListSongGetResp_Privilege{
				Id: item.Sid,
				Maxbr: item.Maxbr,
				St: item.St,
			}

			resp.Privileges = append(resp.Privileges, privilege)
		}
	}
	
	fmt.Println(resp)
	
	return
}