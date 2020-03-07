package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/longhao/music/library/net/http/render"
)


func friendlistGet(c *bm.Context) {
	p := new(struct{
		Uid int64 `json:"uid,omitempty" form:"uid" validate:"required"`
	})

	if err := c.Bind(p); err != nil {
		return
	}

	resp, _ := chatSvc.FriendlistGet(c, p.Uid);
	c.Render(200, render.JSON{resp})
}