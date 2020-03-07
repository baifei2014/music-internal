package model

type Playlist struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	CoverImgUrl string `json:"cover_img_url`
	CreatorUid int64 `json:"creator_uid`
	PlayCount int64 `json:"play_count`
	CommentCount int64 `json:"comment_count"`
	ShareCount int64 `json:"share_count"`
	TrackCount int64 `json:"track_count"`
	SubscribedCount int64 `json:"subscribed_count"`
	Sids []int64 `json:"sids"`
}

func (c Playlist) TableName() string {
	return "playlist"
}

type UserWithPlaylist struct {
	UserId int64
	Pid int64
	Name string `json:"name"`
	CreatorUid int64 `json:"creator_uid`
	CoverImgUrl string `json:"cover_img_url`
	PlayCount int64 `json:"play_count`
	CommentCount int64 `json:"comment_count"`
	ShareCount int64 `json:"share_count"`
	TrackCount int64 `json:"track_count"`
	SubscribedCount int64 `json:"subscribed_count"`
}

func (c UserWithPlaylist) TableName() string {
	return "user_with_playlist"
}

type PlaylistWithSong struct {
	Pid int64 `json:"pid"`
	Sid int64 `json:"sid"`
}

func (c PlaylistWithSong) TableName() string {
	return "playlist_with_song"
}