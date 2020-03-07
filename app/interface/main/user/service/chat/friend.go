package chat

import (
	"fmt"
	"context"

	"github.com/longhao/music/app/interface/main/user/model/chat"

	svcRspb "github.com/longhao/music/app/service/main/music/api/grpc"
)

func (s *Service) FriendlistGet(ctx context.Context, Uid int64) (resp *chat.FriendlistInfoResp, err error) {
	friendlistInfo, err := s.svcCli.FriendlistInfoGet(ctx, &svcRspb.FriendlistInfoGetReq{
		Uid: Uid,
	})

	resp = &chat.FriendlistInfoResp{
		Code: 200,
	}

	var uids []int64
	var userGroup = make(map[int64]string)
	for _, v := range friendlistInfo.List {
		uids = append(uids, v.Uid)
		userGroup[v.Uid] = v.GroupId
	}

	userlistInfo, err := s.dao.BatchUserInfo(ctx, uids)

	fmt.Println("用户信息")
	for _, v := range userlistInfo {
		fmt.Println(v.Nickname)
		elem := &chat.FriendInfo{
			Uid: v.UserId,
			GroupId: userGroup[v.UserId],
			Name: v.Nickname,
			AvatarUrl: v.AvatarUrl,
		}
		resp.Data = append(resp.Data, elem)
	}
	return
}