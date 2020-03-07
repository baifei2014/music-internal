package service

import (
	"fmt"
	"context"

	"github.com/longhao/music/app/service/main/music/dao"
	"github.com/longhao/music/app/service/main/music/conf"
	pb "github.com/longhao/music/app/service/main/music/api/grpc"
)

type FriendResourceService struct {
	dao *dao.Dao
}

func NewFriendResourceService(c *conf.Config) (s *FriendResourceService) {
	s = &FriendResourceService{
		dao: dao.New(c),
	}
	return
}

func (s *FriendResourceService) FriendlistInfoGet(ctx context.Context, req *pb.FriendlistInfoGetReq) (resp *pb.FriendlistInfoGetResp, err error) {
	var (
		uid = req.Uid
		friendUid int64
	)

	friendlist, err := s.dao.GetFriendlistBatch(ctx, uid);


	resp = &pb.FriendlistInfoGetResp{}
	for _, v := range friendlist {
		friendUid = v.ApplyUid
		if friendUid == uid {
			friendUid = v.AcceptUid
		}
		elem := &pb.FriendlistInfoGetResp_List{
			Uid: friendUid,
			GroupId: v.GroupId,
			Name: v.Name,
			AvatarUrl: v.AvatarUrl,
		}

		resp.List = append(resp.List, elem)
	}

	fmt.Println(resp)

	return
}