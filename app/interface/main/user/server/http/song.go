package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/longhao/music/library/net/http/render"
)

type SongResource struct {
	Song *Song `json:"song"`
	Privileges []*Privilege `json:"privileges"`
}

type Song struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Artists []*Artist `json:"artists"`
	Al `json:"al"`
	Duration int32 `json:"duration"`
}

type Artist struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type Al struct {
	Id int32 `json:id`
	Name string `json:name`
	PicUrl string `json:picUrl`
}

type Privilege struct {
	Id int32 `json:"id"`
	Maxbr int32 `json:"maxbr"`
	St int32 `json:"st"`
}

// func songResourceGetDetail(c *bm.Context) {
// 	resp := &SongResource{}

// 	resp.Song = &Song{
// 		Id: 115569,
// 		Name: "护花使者",
// 		Duration: 247000,
// 	}
// 	artist := &Artist {
// 		Id: 3699,
// 		Name: "李克勤",
// 	}
// 	al := Al{
// 		Id: 513122,
// 		Name: "亚洲新生开学了!",
// 		PicUrl: "https://localhost:6006/image/resource/109951164577862384.jpg",
// 	}
// 	resp.Song.Al = al
// 	resp.Song.Artists = append(resp.Song.Artists, artist)

// 	privilege := &Privilege{
// 		Id: 25657384,
// 		Maxbr: 320000,
// 	}
// 	resp.Privileges = append(resp.Privileges, privilege)

// 	c.JSON(http.StatusOK, resp)
// }

func songResourceGetDetail(c *bm.Context) {
	p := new(struct {
		Sid      int64 `json:"sid,omitempty" form:"sid" binding:"required"`
	})
	if err := c.Bind(p); err != nil {
		return
	}

	resp, _ := musicSvc.SongGetDetailWithRelationInfo(c, p.Sid)

	c.JSON(resp, nil)
}

type LyricResource struct {
	Lrc `json:"lrc"`
}

type Lrc struct {
	Lyric string `json:"lyric"`
}

func songLyricGetSingle(c *bm.Context) {
	resp := &LyricResource{
		Lrc: Lrc{
			Lyric: "[by:WLecter]\n[ti:飘摇]\n[ar:周迅]\n[al:]\n[00:15.03]\n[00:21.39]风停了云知道\n[00:25.01]爱走了心自然明了\n[00:29.77]它来时躲不掉\n[00:34.13]它走得静悄悄\n[00:38.52]你不在我预料\n[00:42.21]扰乱我平静的步调\n[00:47.02]怕爱了找苦恼\n[00:51.46]怕不爱睡不着\n[00:57.56]我飘啊飘你摇啊摇\n[01:02.29]无根的野草\n[01:06.62]当梦醒了天晴了\n[01:10.82]如何再飘渺\n[01:14.84]（啊）爱多一秒恨不会少\n[01:19.38]承诺是煎熬\n[01:23.80]若不计较\n[01:25.96]就一次痛快燃烧\n[02:09.15]你不在我预料\n[02:12.70]扰乱我平静的步调\n[02:17.78]怕爱了找苦恼\n[02:22.03]怕不爱睡不着\n[02:28.33]我飘啊飘你摇啊摇\n[02:32.77]无根的野草\n[02:37.15]当梦醒了天晴了\n[02:41.46]如何再飘渺\n[02:45.72]（啊）爱多一秒恨不会少\n[02:50.09]承诺是煎熬\n[02:54.50]若不计较\n[02:56.58]就一次痛快燃烧\n[03:04.90]我飘啊飘你摇啊摇\n[03:09.61]无根的野草\n[03:13.92]当梦醒了天晴了\n[03:18.18]如何再飘渺\n[03:22.53]爱多一秒恨不会少\n[03:26.82]承诺是煎熬\n[03:31.19]若不计较\n[03:33.14]就一次痛快燃烧\n[03:39.77]若不计较\n[03:41.94]就一次痛快燃烧\n",
		},
	}

	c.JSON(resp, nil)
}

func songPlayerUrlInfo(c *bm.Context) {
	p := new(struct{
		Ids []int64 `json:"ids,omitempty" form:"ids" validate:"required"`
		Br int64 `json:"br,omitempty" form:"br"`
	})

	if err := c.Bind(p); err != nil {
		return
	}

	resp, _ := musicSvc.SongListInfoGet(c, p.Ids);
	c.Render(200, render.JSON{resp})
}