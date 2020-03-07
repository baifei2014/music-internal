package dao

import (
	"context"
	"github.com/longhao/music/app/interface/main/user/model"
	rspb "github.com/longhao/music/app/service/main/music/api/grpc"
	"github.com/longhao/music/library/xstr"
)

func (d *Dao) BatchUserInfo(ctx context.Context, uids []int64) (resp map[int64]*model.UserBase, err error) {
	userInfo, err := d.userClient.ListUserInfo(ctx, &rspb.ListUserInfoReq{
		Uid: uids,
	})

	resp = make(map[int64]*model.UserBase)
	for _, v := range userInfo.List {
		elem := &model.UserBase{
			UserId: v.Uid,
			Nickname: v.Nickname,
			AvatarUrl: xstr.GetFullImageUrl(v.AvatarUrl),
		}
		resp[v.Uid] = elem
	}

	return
}