package v1

import (
	"google.golang.org/grpc"
)

type Client struct {
	SongResourceClient
	PlaylistResourceClient
	UserResourceClient
	FriendResourceClient
	MessageResourceClient
}

const (
	address     = "localhost:50052"
)

func NewClient() (*Client) {
	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	cli := &Client{}
	cli.SongResourceClient = NewSongResourceClient(conn)
	cli.PlaylistResourceClient = NewPlaylistResourceClient(conn)
	cli.UserResourceClient = NewUserResourceClient(conn)
	cli.FriendResourceClient = NewFriendResourceClient(conn)
	cli.MessageResourceClient = NewMessageResourceClient(conn)
	return cli
}