package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"

	"github.com/longhao/music/library/net/http/render"
)

type CateResource struct {
	All `json:"all"`
}

type All struct {
	Name string `json:"name"`
}

func cateListGet(c *bm.Context) {
	resp := &CateResource{
		All: All{
			Name: "全部歌单",
		},
	}

	c.JSON(resp, nil)
}

type PlayListResource struct {
	PlayLists []*PlayList `json:"playlists"`
}

type PlayList struct  {
	Id int64 `json:"id"`
	Name string `json:"name"`
	CoverImgUrl string `json:"coverImgUrl"`
	Tracks []*Track `json:"tracks"`
	Creator `json:"creator"`
	PlayCount int32 `json:"playCount"`
	CommentCount int32 `json:"commentCount"`
	ShareCount int32 `json:"shareCount"`
}

type Track struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Ar []*Artist `json:"ar"`
	Al *Album `json:"al"`
}

type Creator struct {
	UserId int32 `json:"userId"`
	AvatarUrl string `json:"avatarUrl`
	Nickname string `json:"nickname"`
}

type Album struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	PicUrl string `json:"picUrl"`
}




func playListGet(c *bm.Context) {
	resp, _ := musicSvc.PlaylistGet(c)


	c.JSON(resp, nil)
}

type PlayListDetailResource struct {
	PlayList *PlayList `json:"playlist"`
	Privileges []*Privilege `json:"privileges"`
}


func playListGetDetail(c *bm.Context) {
	p := new(struct{
		Id int64 `json:"id,omitempty" form:"id" binding:"required"`
	})

	if err := c.Bind(p); err != nil {
		return
	}

	resp, _ := musicSvc.PlaylistGetDetail(c, p.Id)

	c.JSON(resp, nil)
}

func playlistDetailInfo(c *bm.Context) {
	p := new(struct{
		Id int64 `json:"id,omitempty" form:"id" validate:"required"`
	})

	if err := c.Bind(p); err != nil {
		return
	}

	resp, _ := musicSvc.GetPlaylistDetailInfo(c, p.Id)

	c.Render(200, render.JSON{resp})
}