package dao

import (
	"fmt"
	"context"

	"github.com/longhao/music/app/service/main/music/model"
)

func (d *Dao) GetMessageListBatch(ctx context.Context, groupId string, page int64, pageSize int64) (list []model.Message, err error) {
	err = d.rsDBReader.Model(&model.Message{}).
		Where("group_id = ?", groupId).
		Select("*").
		Offset((page-1)*pageSize).
		Limit(pageSize).
		Find(&list).Error
	fmt.Println(err)

	fmt.Println(list)
	return
}