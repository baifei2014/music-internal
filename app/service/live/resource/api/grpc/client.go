package v1

import (
	"google.golang.org/grpc"
)

type Client struct {
	UserResourceClient
	BannerResourceClient
}

const (
	address     = "localhost:50051"
)

func NewClient() (*Client) {
	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	cli := &Client{}
	cli.UserResourceClient = NewUserResourceClient(conn)
	cli.BannerResourceClient = NewBannerResourceClient(conn)
	return cli
}