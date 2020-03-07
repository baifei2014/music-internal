package model

type UserResourceGetSingleResp struct {
	Uid                  int32    `json:"uid"`
	Username             string   `json:"username"`
}

type UserBase struct {
	UserId int64 `json:"userId"`
	Nickname string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
}