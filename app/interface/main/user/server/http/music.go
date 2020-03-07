
package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"

)

type UrlResource struct {
	Data []*MusicUrl `json:"data"`
}

type MusicUrl struct {
	Url string `json:"url"`
}

// func musicUrlGet(c *gin.Context) {
// 	resp := &UrlResource{}

// 	musicUrl := &MusicUrl{
// 		Url: "http://localhost:6006/song/070e_0e0e_010e_0d80a760bc9d091e9eee6d1a322dd0ed.m4a",
// 	}
// 	resp.Data = append(resp.Data, musicUrl)

// 	c.JSON(http.StatusOK, resp)
// }

func musicUrlGet(c *bm.Context) {
	p := new(struct {
		Sid      int64 `json:"sid,omitempty" form:"sid" binding:"required"`
	})
	if err := c.Bind(p); err != nil {
		return
	}

	resp, _ := musicSvc.SongGetDetail(c, p.Sid)

	c.JSON(resp, nil)
}