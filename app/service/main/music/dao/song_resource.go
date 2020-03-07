package dao

import (
	"fmt"
	"context"
	"database/sql"
	// "time"

	"github.com/longhao/music/app/service/main/music/model"

	"github.com/jinzhu/gorm"
)

var (
	rowFields = "`id`, `sid`,`name`, `duration`, `url`, `is_delete`,`create_time`,`update_time`"
)

// getRowResult Helper方法
func getRowResult(queryResult *gorm.DB) (res model.SongResource, err error) {
	var count int64

	err = queryResult.Count(&count).Error
	if err != nil {
		fmt.Println("获取总数查询出错")
		return
	}

	if count == 0 {
		fmt.Println("获取总数为空查询出错")
		return
	}

	var id, sid, duration, isDelete sql.NullInt64
	var name sql.NullString
	var createTime, updateTime sql.NullTime

	err = queryResult.Row().Scan(&id, &sid, &name, &duration, &isDelete, &createTime, &updateTime)

	if err != nil {
		fmt.Println("分子端查询出错")
		panic(err)
		return
	}

	res.Id = int64(id.Int64)
	res.Sid = int64(sid.Int64)
	res.Name = string(name.String)
	res.Duration = int64(duration.Int64)
	res.IsDelete = int8(isDelete.Int64)
	res.CreateTime = createTime.Time
	res.UpdateTime = updateTime.Time

	return
}

func (d *Dao) GetPlaylist(ctx context.Context) (resp []model.Playlist, err error) {
	d.rsDBReader.Model(&model.Playlist{}).Select("*").Find(&resp)

	return 
}

func (d *Dao) GetPrivilegeInfo(ctx context.Context, sids []int64) (res []model.Privilege, err error) {
	d.rsDBReader.Model(&model.Privilege{}).
		Where("sid IN (?)", sids).
		Select("privilege.*").
		Find(&res)

	return
}

func (d *Dao) GetAlbumInfo(ctx context.Context, sids []int64) (res []*model.Album, err error) {
	var albumItem []model.SongWithAlbum
	d.rsDBReader.Model(&model.SongWithAlbum{}).
		Where("song_with_album.sid IN (?)", sids).
		Joins("left join album on song_with_album.tid = album.id").
		Select("song_with_album.*, album.name as name, album.pic_url as pic_url").
		Find(&albumItem)

	for _, item := range albumItem {
		album := &model.Album{
			Id: item.Tid,
			Sid: item.Sid,
			Name: item.Name,
			PicUrl: item.PicUrl,
		}
		res = append(res, album)
	}

	return
}

func (d *Dao) GetArtistInfo(ctx context.Context, sids []int64) (res []model.Artist, err error) {
	var songWithArtistList []model.SongWithArtist
	d.rsDBReader.Model(&model.SongWithArtist{}).
		Where("song_with_artist.sid in (?)", sids).
		Joins("left join artist on song_with_artist.tid = artist.id").
		Select("song_with_artist.*, artist.name as name").
		Find(&songWithArtistList)

	for _, item := range songWithArtistList {
		ar := model.Artist{
			Id: item.Tid,
			Sid: item.Sid,
			Name: item.Name,
		}
		res = append(res, ar)
	}

	return 
}

func (d *Dao) GetSongResourceInfo(ctx context.Context, sids []int64) (res []model.SongResource, err error) {
	d.rsDBReader.Model(&model.SongResource{}).Select(rowFields).
		Where("sid IN (?)", sids).
		Find(&res)

	return
}