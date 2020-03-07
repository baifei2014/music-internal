package music

import (
	"github.com/longhao/music/app/interface/main/user/model"
)

type PlaylistDetailResp struct {
	Code int64 `json:"code"`
	Playlist *Playlist `json:"playlist"`
	Privileges []*Privilege `json:"privileges"`
}

type Playlist struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	CoverImgUrl string `json:"coverImgUrl"`
	Creator *model.UserBase `json:"creator"`
	PlayCount int64 `json:"playCount"`
	CommentCount int64 `json:"commentCount"`
	ShareCount int64 `json:"shareCount"`
	TrackCount int64 `json:"trackCount"`
	Tracks []*Song `json:"tracks"`
	SubscribedCount int64 `json:"subscribedCount"`
	Subscribed bool `json:"subscribed"`
}

type PlaylistResource struct {
	Playlists []*PlaylistResp `json:"playlists"`
}

type PlaylistResp struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	CoverImgUrl string `json:"coverImgUrl"`
}

type UserPlaylistResp struct {
	More bool `json:"more"`
	Playlist []*Playlist `json:"playlist"`
	Code int32 `json:"code"`
}