package dao

import (
	"fmt"
	"context"

	"github.com/longhao/music/app/service/main/music/model"
)

func(d *Dao) GetFriendlistBatch(ctx context.Context, Uid int64) (list []model.Friendlist, err error) {
	err = d.rsDBReader.Model(&model.Friendlist{}).
		Where("apply_uid = ?", Uid).
		Where("status = ?", 1).
		Select("*").
		Find(&list).Error
	fmt.Println(err)

	fmt.Println(list)
	return
}