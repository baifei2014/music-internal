package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func bannerResourceGet(c *bm.Context) {
	p := new(struct {
		Page int32 `json:"page,omitempty" form:"page"`
		PageSize int32 `json:"page,omitempty" form:"page_size"`
	})

	if err := c.Bind(p); err != nil {
		return
	}
	resp, _ := webSvc.GetBanner(c, p.Page, p.PageSize)

	c.JSON(resp, nil)
}