package service

import (
	"fmt"
	"context"

	"github.com/longhao/music/app/service/main/music/dao"
	"github.com/longhao/music/app/service/main/music/conf"
	pb "github.com/longhao/music/app/service/main/music/api/grpc"
)

type MessageResourceService struct {
	dao *dao.Dao
}

func NewMessageResourceService(c *conf.Config) (s *MessageResourceService) {
	s = &MessageResourceService{
		dao: dao.New(c),
	}
	return
}

func (s *MessageResourceService) MessageList(ctx context.Context, req *pb.MessageListReq) (resp *pb.MessageListResp, err error) {
	var (
		Page int64 = 1
		PageSize int64 = 50
	)

	if req.Page > 0 {
		Page = req.Page
	}
	if req.PageSize > 0 {
		PageSize = req.PageSize
	}

	resp = &pb.MessageListResp{
		CurrentPage: Page,
		More: false,
	}

	messageList, err := s.dao.GetMessageListBatch(ctx, req.GroupId, Page, PageSize)

	if len(messageList) <= 0 {
		return
	}

	resp.More = true
	for _, v := range messageList{
		elem := &pb.MessageListResp_List{
			FromUid: v.FromUid,
			ToUid: v.ToUid,
			Content: v.Content,
		}
		resp.List = append(resp.List, elem)
	}

	fmt.Println(resp)
	return
}