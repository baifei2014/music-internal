package chat

type FriendlistInfoResp struct {
	Code int64 `json:"code"`
	Data []*FriendInfo `json:"data"`
}

type FriendInfo struct {
	Uid int64 `json:"uid"`
	GroupId string `json:"groupId"`
	Name string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}