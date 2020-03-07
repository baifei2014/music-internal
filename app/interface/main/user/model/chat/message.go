package chat

type MessageListResp struct {
	Code int64 `json:"code"`
	CurrentPage int64 `json:"current_pagee"`
	More bool `json:"more"`
	MessageList []*Message `json:"message_list"`
}

type Message struct {
	Uid int64 `json:"uid"`
	Content string `json:"content"`
	AvatarUrl string `json:"avatarUrl"`
	IsSelf bool `json:"is_self"`
}