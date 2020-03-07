package music

import (
	"fmt"
	"context"

	"github.com/longhao/music/app/interface/main/user/model/music"

	rspb "github.com/longhao/music/app/service/main/music/api/grpc"

	"github.com/longhao/music/library/xstr"
)

func (s *Service) PlaylistGet(ctx context.Context) (resp *music.PlaylistResource, err error) {
	respRPC, err := s.svcCli.PlaylistGetList(ctx, &rspb.PlaylistGetListReq{})

	resp = &music.PlaylistResource{}
	for _, item := range respRPC.Playlists {
		playlist := &music.PlaylistResp{
			Id: item.Id,
			Name: item.Name,
			CoverImgUrl: item.CoverImgUrl,
		}
		resp.Playlists = append(resp.Playlists, playlist)
	}

	return
}

func (s *Service) PlaylistGetDetail(ctx context.Context, id int64) (resp *music.PlaylistDetailResp, err error) {
	respRPC, err := s.svcCli.PlaylistGetSingle(ctx, &rspb.GetPlaylistSingleReq{
		Id: id,
	})

	fmt.Println(respRPC);

	resp = &music.PlaylistDetailResp{}

	resp.Playlist = &music.Playlist{
		Id: respRPC.Id,
		Name: respRPC.Name,
		CoverImgUrl: xstr.GetFullImageUrl(respRPC.CoverImgUrl),
		PlayCount: respRPC.PlayCount,
		CommentCount: respRPC.CommentCount,
		ShareCount: respRPC.ShareCount,
	}

	songGetListRespRPC, err := s.svcCli.SongGetDetail(ctx, &rspb.GetDetailReq{
		Sids: respRPC.Sids,
	})

	albumRespRPC, _ := s.svcCli.AlbumGet(ctx, &rspb.GetAlbumReq{
		Sids: respRPC.Sids,
	})

	artistRespRPC, _ := s.svcCli.ArtistGet(ctx, &rspb.GetArtistReq{
		Sids: respRPC.Sids,
	})

	artists := make(map[int64]*music.ArtistGroup)
	for _, item := range artistRespRPC.Artists {
		if artists[item.Sid] == nil {
			artists[item.Sid] = &music.ArtistGroup{}
		}
		artist := &music.Artist{
			Id: item.Id,
			Name: item.Name,
		}
		artists[item.Sid].Artists = append(artists[item.Sid].Artists, artist)
	}

	albums := make(map[int64]*music.AlbumGroup)
	for _, item := range albumRespRPC.Albums {
		if albums[item.Sid] == nil {
			albums[item.Sid] = &music.AlbumGroup{}
		}
		al := &music.Album{
			Id: item.Id,
			Name: item.Name,
			PicUrl: xstr.GetFullImageUrl(item.PicUrl),
		}
		albums[item.Sid].Album = al
	}

	for _, songItem := range songGetListRespRPC.Songs {
		track := &music.Song{
			Id: songItem.Id,
			Name: songItem.Name,
			Duration: songItem.Duration,
			Al: albums[songItem.Id].Album,
			Artists: artists[songItem.Id].Artists,
		}

		resp.Playlist.Tracks = append(resp.Playlist.Tracks, track)
	}

	for _, privilegeItem := range songGetListRespRPC.Privileges {
		privilege := &music.Privilege{
			Id: privilegeItem.Id,
			Maxbr: privilegeItem.Maxbr,
			St: privilegeItem.St,
		}
		resp.Privileges = append(resp.Privileges, privilege)
	}

	fmt.Println(resp)

	return
}

func (s *Service) GetPlaylistDetailInfo(ctx context.Context, id int64) (resp *music.PlaylistDetailResp, err error) {
	respRPC, err := s.svcCli.PlaylistDetailGet(ctx, &rspb.PlaylistDetailGetReq{
		Id: id,
	})

	fmt.Println("远程请求接口")
	fmt.Println(respRPC)

	resp = &music.PlaylistDetailResp{
		Code: 200,
	}

	var userIds = []int64{respRPC.CreatorUid}
	userInfos, _ := s.batchUserInfo(ctx, userIds)

	resp.Playlist = &music.Playlist{
		Id: respRPC.Id,
		Name: respRPC.Name,
		Creator: userInfos[respRPC.CreatorUid],
		CoverImgUrl: xstr.GetFullImageUrl(respRPC.CoverImgUrl),
		PlayCount: respRPC.PlayCount,
		CommentCount: respRPC.CommentCount,
		ShareCount: respRPC.ShareCount,
		SubscribedCount: respRPC.SubscribedCount,
		Subscribed: true,
	}

	songGetListRespRPC, err := s.svcCli.ListSongGet(ctx, &rspb.ListSongGetReq{
		Sids: respRPC.Sids,
	})

	albumRespRPC, _ := s.svcCli.AlbumGet(ctx, &rspb.GetAlbumReq{
		Sids: respRPC.Sids,
	})

	artistRespRPC, _ := s.svcCli.ArtistGet(ctx, &rspb.GetArtistReq{
		Sids: respRPC.Sids,
	})

	artists := make(map[int64]*music.ArtistGroup)
	for _, item := range artistRespRPC.Artists {
		if artists[item.Sid] == nil {
			artists[item.Sid] = &music.ArtistGroup{}
		}
		artist := &music.Artist{
			Id: item.Id,
			Name: item.Name,
		}
		artists[item.Sid].Artists = append(artists[item.Sid].Artists, artist)
	}

	albums := make(map[int64]*music.AlbumGroup)
	for _, item := range albumRespRPC.Albums {
		if albums[item.Sid] == nil {
			albums[item.Sid] = &music.AlbumGroup{}
		}
		al := &music.Album{
			Id: item.Id,
			Name: item.Name,
			PicUrl: xstr.GetFullImageUrl(item.PicUrl),
		}
		albums[item.Sid].Album = al
	}

	for _, songItem := range songGetListRespRPC.Songs {
		track := &music.Song{
			Id: songItem.Id,
			Name: songItem.Name,
			Duration: songItem.Duration,
			Url: xstr.GetFullImageUrl(songItem.Url),
			Al: albums[songItem.Id].Album,
			Artists: artists[songItem.Id].Artists,
		}

		resp.Playlist.Tracks = append(resp.Playlist.Tracks, track)
	}

	for _, privilegeItem := range songGetListRespRPC.Privileges {
		privilege := &music.Privilege{
			Id: privilegeItem.Id,
			Maxbr: privilegeItem.Maxbr,
			St: privilegeItem.St,
		}
		resp.Privileges = append(resp.Privileges, privilege)
	}

	fmt.Println(resp)

	return
}