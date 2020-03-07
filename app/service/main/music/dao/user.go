package dao

import (
	"context"
	"github.com/longhao/music/app/service/main/music/model"
)

func (d *Dao) BatchUserInfo(ctx context.Context, uids []int64) (list []*model.User, err error) {
	d.rsDBReader.Model(&model.User{}).
		Where("uid IN (?)", uids).
		Select("*").
		Find(&list)

	return
}