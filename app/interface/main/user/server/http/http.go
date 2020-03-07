package http

import (
	// "time"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/longhao/music/app/interface/main/user/service"
	"github.com/longhao/music/app/interface/main/user/service/music"
	"github.com/longhao/music/app/interface/main/user/service/chat"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

var (
	webSvc *service.Service
	musicSvc *music.Service
	chatSvc *chat.Service
)

func Init() {
	// engine := gin.Default()
	var timeout xtime.Duration
	timeout.UnmarshalText([]byte("1s"))
	serverConfig := &bm.ServerConfig{
		Network: "tcp",
		Addr: "0.0.0.0:8080",
		Timeout: timeout,
		ReadTimeout: timeout,
		WriteTimeout: timeout,
	}
	engine := bm.DefaultServer(serverConfig);

	webSvc = service.New()
	musicSvc = music.New()
	chatSvc = chat.New()
	wxroute(engine)
	flutterRoute(engine)

	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func wxroute(e *bm.Engine) {
	group := e.Group("/api")
	{
		userGroup := group.Group("/user")
		{
			userGroup.GET("/info", userResourceGetSingle)
			userGroup.GET("/playlist", userPlayGet)
		}

		bannerGroup := group.Group("banner")
		{
			bannerGroup.GET("/list", bannerResourceGet)
		}

		playGroup := group.Group("playlist")
		{
			playGroup.GET("/catlist", cateListGet)
			playGroup.GET("/list", playListGet)
			playGroup.GET("/detail", playListGetDetail)
		}

		songGroup := group.Group("song")
		{
			songGroup.GET("/detail", songResourceGetDetail)
			songGroup.GET("/lyric", songLyricGetSingle)
		}

		musicGroup := group.Group("music")
		{
			musicGroup.GET("/url", musicUrlGet)
		}

		commentGroup := group.Group("comment")
		{
			commentGroup.GET("/get", commentGet)
		}

	}
}

func flutterRoute(e *bm.Engine) {
	group := e.Group("weapi")
	{
		userGroup := group.Group("user")
		{
			userGroup.POST("playlist", getUserPlayList)
		}
		playlistGroup := group.Group("playlist")
		{
			playlistGroup.POST("detail", playlistDetailInfo)
		}
		songGroup := group.Group("song")
		{
			songGroup.POST("enhance/player/url", songPlayerUrlInfo)
		}
		friendGroup := group.Group("friend")
		{
			friendGroup.POST("list", friendlistGet)
		}
		messageGroup := group.Group("chat_message")
		{
			messageGroup.POST("list", getMessageListInfo)
		}
	}
}