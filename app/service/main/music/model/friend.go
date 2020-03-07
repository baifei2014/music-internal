package model

type Friendlist struct {
	ApplyUid int64 `json:"apply_uid"`
	AcceptUid int64 `json:"accept_uid"`
	GroupId string `json:"group_id"`
	Name string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

func (c Friendlist) TableName() string {
	return "friend"
}