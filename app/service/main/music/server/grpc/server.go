package grpc

import (
	"net"
	pb "github.com/longhao/music/app/service/main/music/api/grpc"
	"github.com/longhao/music/app/service/main/music/conf"
	svc "github.com/longhao/music/app/service/main/music/service"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

// New
func New(c *conf.Config) *grpc.Server {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterSongResourceServer(s, svc.NewSongResourceService(c))
	pb.RegisterPlaylistResourceServer(s, svc.NewPlaylistResourceService(c))
	pb.RegisterUserResourceServer(s, svc.NewUserResourceService(c))
	pb.RegisterFriendResourceServer(s, svc.NewFriendResourceService(c))
	pb.RegisterMessageResourceServer(s, svc.NewMessageResourceService(c))
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
	return s
}
