package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/longhao/music/library/net/http/render"
)


func getMessageListInfo(c *bm.Context) {
	p := new(struct{
		GroupId string `json:"group_id,omitempty" form:"group_id" validate:"required"`
		Page int64 `json:"page,omitempty" form:"page" validate:"required"`
		PageSize int64 `json:"page_size,omitempty" form:"page_size" validate:"required"`
	})

	if err := c.Bind(p); err != nil {
		return
	}

	resp, _ := chatSvc.MessageListGet(c, p.GroupId, p.Page, p.PageSize);
	
	c.Render(200, render.JSON{resp})
}