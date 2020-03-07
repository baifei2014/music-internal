package dao

import (
	"fmt"
	"context"
	"database/sql"
	// "time"

	"github.com/longhao/music/app/service/live/resource/model"

	"github.com/jinzhu/gorm"
)

var (
	rowFields = "`id`, `uid`,`username`,`is_delete`,`create_time`,`update_time`"
)

// getRowResult Helper方法
func getRowResult(queryResult *gorm.DB) (res model.UserResource, err error) {
	var count int32

	err = queryResult.Count(&count).Error
	if err != nil {
		fmt.Println("获取总数查询出错")
		return
	}

	if count == 0 {
		fmt.Println("获取总数为空查询出错")
		return
	}

	var id, uid, isDelete sql.NullInt64
	var username sql.NullString
	var createTime, updateTime sql.NullTime

	err = queryResult.Row().Scan(&id, &uid, &username, &isDelete, &createTime, &updateTime)

	if err != nil {
		fmt.Println("分子端查询出错")
		panic(err)
		return
	}

	res.ID = int32(id.Int64)
	res.Uid = int32(uid.Int64)
	res.Username = string(username.String)
	res.IsDelete = int8(isDelete.Int64)
	res.CreateTime = createTime.Time
	res.UpdateTime = updateTime.Time

	return
}

func (d *Dao) GetUserResourceInfo(ctx context.Context, uid int32) (res model.UserResource, err error) {
	tableInfo := &model.UserResource{}

	queryResult := d.rsDBReader.Model(tableInfo).Select(rowFields).
		Where("uid=?", uid)

	res, err = getRowResult(queryResult)
	if err != nil {
		fmt.Println("查询出错")
	}

	return
}