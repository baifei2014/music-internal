package music

import (
	"fmt"
	"context"

	"github.com/longhao/music/app/interface/main/user/model/music"

	rspb "github.com/longhao/music/app/service/main/music/api/grpc"
	"github.com/longhao/music/library/xstr"
)


func (s *Service) SongListInfoGet(ctx context.Context, sids []int64) (resp *music.SongPlayerInfoResp, err error) {
	songGetListRespRPC, err := s.svcCli.ListSongGet(ctx, &rspb.ListSongGetReq{
		Sids: sids,
	})

	fmt.Println(songGetListRespRPC)

	resp = &music.SongPlayerInfoResp{
		Code: 200,
	}
	for _, v := range songGetListRespRPC.Songs {
		elem := &music.SongPlayerInfo{
			Id: v.Id,
			Name: v.Name,
			Duration: v.Duration,
			Url: xstr.GetFullImageUrl(v.Url),
		}
		resp.Data = append(resp.Data, elem)
	}

	return
}

func (s *Service) SongGetDetail(ctx context.Context, sid int64) (resp *music.PlaySongResp, err error) {
	respRPC, err := s.svcCli.SongGetDetail(ctx, &rspb.GetDetailReq{
		Sids: []int64{sid},
	})


	resp = &music.PlaySongResp{}
	for _, item := range respRPC.Songs {
		song := &music.PlaySong{
			Id: item.Id,
			Name: item.Name,
			Duration: item.Duration,
			Url: item.Url,
		}
		resp.Data = append(resp.Data, song)
	}

	fmt.Println(resp)

	return
}

func (s *Service) SongGetDetailWithRelationInfo(ctx context.Context, sid int64) (resp *music.SongResourceGetDetailResp, err error) {
	respRPC, err := s.svcCli.SongGetDetail(ctx, &rspb.GetDetailReq{
		Sids: []int64{sid},
	})


	if err == nil {
		resp = &music.SongResourceGetDetailResp{
			Songs: convertRPCSongRes(respRPC.Songs),
			Privileges: convertRPCPrivilegesRes(respRPC.Privileges),
		}
	}

	albumRespRPC, _ := s.svcCli.AlbumGet(ctx, &rspb.GetAlbumReq{
		Sids: []int64{sid},
	})

	artistRespRPC, _ := s.svcCli.ArtistGet(ctx, &rspb.GetArtistReq{
		Sids: []int64{sid},
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
			PicUrl: item.PicUrl,
		}
		albums[item.Sid].Album = al
	}

	for index, item := range resp.Songs {
		resp.Songs[index].Al = albums[item.Id].Album
		resp.Songs[index].Artists = artists[item.Id].Artists
	}

	fmt.Println(resp)

	return
}

func convertRPCPrivilegesRes(RPCPrivileges []*rspb.GetDetailResp_Privilege) (HTTPPrivileges []*music.Privilege) {
	HTTPPrivileges = make([]*music.Privilege, len(RPCPrivileges))

	for index, privilegeItem := range RPCPrivileges {
		HTTPPrivileges[index] = &music.Privilege{
			Id: privilegeItem.Id,
			Maxbr: privilegeItem.Maxbr,
			St: privilegeItem.St,
		}
	}

	return
}

func convertRPCSongRes(RPCSong []*rspb.GetDetailResp_Song) (HTTPSong []*music.Song) {
	for _, item := range RPCSong {
		song := &music.Song{
			Id: item.Id,
			Name: item.Name,
			Duration: item.Duration,
		}
		HTTPSong = append(HTTPSong, song)
	}

	return
}

