package model

type User struct {
	Uid int64 `json:"uid"`
	Nickname string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
}

func (c User) TableName() string {
	return "user"
}