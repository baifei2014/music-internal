package model

import "time"

type SongResource struct {
	Id int64 `json:"id"`
	Sid int64 `json:"sid"`
	Name string `json:"name"`
	Duration int64 `json:"duration"`
	Url string `json:"url"`
	IsDelete int8 `json:"is_delete"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (c SongResource) TableName() string {
	return "song"
}

type SongWithAlbum struct {
	Sid int64 `json:"sid" gorm:"column:sid"`
	Tid int64 `json:"tid" gorm:"column:tid"`
	Name string `json:"name" gorm:"-"`
	PicUrl string `json:"pic_url" gorm:"-"`
}

func (c SongWithAlbum) TableName() string {
	return "song_with_album"
}

type Album struct {
	Id int64 `json:"id"`
	Sid int64 `json:"sid"`
	Name string `json:"name"`
	PicUrl string `json:"pic_url"`
	IsDelete int8 `json:"is_delete"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (c Album) TableName() string {
	return "album"
}

type SongWithArtist struct {
	Sid int64 `json:"sid" gorm:"column:sid"`
	Tid int64 `json:"tid" gorm:"column:tid"`
	Name string `json:"name" gorm:"-"`
}

func (c SongWithArtist) TableName() string {
	return "song_with_artist"
}

type Artist struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Sid int64 `json:"sid"`
	IsDelete int8 `json:"is_delete"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (c Artist) TableName() string {
	return "artist"
}

type Privilege struct {
	Sid int64 `json:"mid"`
	Maxbr int64 `json:"maxbr"`
	St int64 `json:"st"`
}

func (c Privilege) TableName() string {
	return "privilege"
}