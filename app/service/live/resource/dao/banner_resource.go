package dao

import (
	"context"
	"database/sql"

	"github.com/longhao/music/app/service/live/resource/model"
)

func (d *Dao) GetTotalCount(ctx context.Context) (totalCount int32, err error) {
	tableInfo := &model.BannerResource{}

	var ret sql.NullInt64

	err = d.rsDB.Model(tableInfo).Debug().
		Select("count(*) as total").
		Row().Scan(&ret)

	if err != nil {
		return
	}

	totalCount = int32(ret.Int64)

	return
}

func (d *Dao) ListBannerResourceInfo(ctx context.Context, page int32, pageSize int32) (list []model.BannerResource, err error) {
	var tx = d.rsDBReader
	tableInfo := &model.BannerResource{}

	err = tx.Model(tableInfo).
		Select("`id`, `pic`, `is_delete`, `create_time`, `update_time`").
		Order("id ASC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&list).Error

	if err != nil {
		return
	}

	return
}