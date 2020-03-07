package chat

import (
	"fmt"
	"context"

	"github.com/longhao/music/app/interface/main/user/model/chat"

	svcRspb "github.com/longhao/music/app/service/main/music/api/grpc"
)

func (s *Service) MessageListGet(ctx context.Context, groupId string, page int64, pageSize int64) (resp *chat.MessageListResp, err error) {
	messageList, err := s.svcCli.MessageList(ctx, &svcRspb.MessageListReq{
		GroupId: groupId,
		Page: page,
		PageSize: pageSize,
	})

	uids := []int64{}

	for _, v := range messageList.List {
		uids = append(uids, v.FromUid , v.ToUid)
	}

	userlistInfo, err := s.dao.BatchUserInfo(ctx, uids)

	resp = &chat.MessageListResp{
		Code: 200,
		CurrentPage: messageList.CurrentPage,
		More: messageList.More,
		MessageList: []*chat.Message{},
	}
	for _, v := range messageList.List {
		elem := &chat.Message{
			Uid: v.FromUid,
			Content: v.Content,
			AvatarUrl: userlistInfo[v.FromUid].AvatarUrl,
			IsSelf: v.FromUid == 97571233,
		}
		resp.MessageList = append(resp.MessageList, elem)
	}

	fmt.Println(resp)

	return
}