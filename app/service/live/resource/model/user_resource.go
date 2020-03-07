package model

import "time"

type UserResource struct {
	ID int32 `json:"id"`
	Uid int32 `json:"uid"`
	Username string `json:"username"`
	IsDelete int8 `json:"is_delete"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (c UserResource) TableName() string {
	return "user_resource"
}