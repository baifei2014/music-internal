package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"

	"github.com/longhao/music/library/net/http/render"
)

type User struct {
	Account *Account `json:"account"`
}

type Account struct {
	Id int32 `json:"id"`
	Username string `json:"username"`
}

func userResourceGetSingle(c *bm.Context) {
	// p := new(struct {
	// 	Uid      int32 `json:"uid,omitempty" form:"uid" validate:"required"`
	// })
	// if err := c.ShouldBind(p); err != nil {
	// 	return
	// }
	// resp, _ := webSvc.GetUserSingle(c, p.Uid)
	
	resp := &User{}

	resp.Account = &Account{
		Id: 97571233,
		Username: "1_15237391071",
	}
	
	c.JSON(resp, nil)
}

type UserPlay struct {
	PlayList []*Play `json:"playlist"`
}

type Play struct {
	UserId int32 `json:"userId"`
}

func userPlayGet(c *bm.Context) {
	resp := &UserPlay{}

	play := &Play{
		UserId: 97571233,
	}
	resp.PlayList = append(resp.PlayList, play)
	
	c.JSON(resp, nil)
}

func getUserPlayList(c *bm.Context) {
	p := new(struct{
		Uid int64 `json:"uid,omitempty" form:"uid" validate:"required"`
		Limit int64 `json:"limit,omitempty" form:"limit" validate:"required"`
		Page int64 `json:"page,omitempty" form:"page" validate:"required"`
	})

	if err := c.Bind(p); err != nil {
		return
	}

	resp, _ := musicSvc.UserPlaylistGet(c, p.Uid, p.Page, p.Limit)

	c.Render(200, render.JSON{resp})
}