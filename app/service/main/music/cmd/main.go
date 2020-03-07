package main

import (
	// "os"
	// "syscall"
	// "time"
	// "os/signal"
	"flag"
	// "fmt"

	"github.com/longhao/music/app/service/main/music/conf"
	"github.com/longhao/music/app/service/main/music/server/grpc"

)

func main() {
	flag.Parse()

	if err := conf.Init(); err != nil {
		panic(err)
	}

	grpc.New(conf.Conf)

	// http.Init()
	// c := make(chan os.Signal, 1)
	// signal.Notify(c)

	// for {
	// 	s := <-c
	// 	log.Info("get a signal %s", s.String())
	// 	switch s {
	// 	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
	// 		log.Info("live-admin-admin exit")
	// 		time.Sleep(time.Second)
	// 		return
	// 	case syscall.SIGHUP:
	// 	default:
	// 		return
	// 	}
	// }
}