package model

type Message struct {
	FromUid int64 `json:"from_uid"`
	ToUid int64 `json:"to_uid"`
	Content string `json:"content"`
}

func (c Message) TableName() string {
	return "message"
}