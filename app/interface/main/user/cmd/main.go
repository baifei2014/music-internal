package main

import (
	"os"
	"syscall"
	"time"
	"os/signal"

	"github.com/longhao/music/app/interface/main/user/server/http"
	"github.com/longhao/music/library/log"
)

func main() {
	http.Init()
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("live-admin-admin exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}