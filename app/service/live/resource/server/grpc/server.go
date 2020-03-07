package grpc

import (
	"net"
	pb "github.com/longhao/music/app/service/live/resource/api/grpc"
	"github.com/longhao/music/app/service/live/resource/conf"
	svc "github.com/longhao/music/app/service/live/resource/service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// New
func New(c *conf.Config) *grpc.Server {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserResourceServer(s, svc.NewUserResourceService(c))
	pb.RegisterBannerResourceServer(s, svc.NewBannerResourceService(c))
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
	return s
}
