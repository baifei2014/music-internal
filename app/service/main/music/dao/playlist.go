package dao

import (
	"fmt"
	"context"
	"database/sql"

	"github.com/longhao/music/app/service/main/music/model"
)

func (d *Dao) GetTotalCount(ctx context.Context, userId int64) (totalCount int64, err error) {
	var ret sql.NullInt64

	err = d.rsDB.Model(&model.UserWithPlaylist{}).Debug().
		Where("user_id =?", userId).
		Select("count(*) as total").
		Row().Scan(&ret)

	if err != nil {
		return
	}

	totalCount = int64(ret.Int64)

	return
}

func (d *Dao) ListUserPlaylistResourceInfo(ctx context.Context, userId int64, page int64, pageSize int64) (list []*model.Playlist, err error) {
	var (
		listwith []model.UserWithPlaylist
	)
	err = d.rsDBReader.Model(&model.UserWithPlaylist{}).
		Joins("left join playlist on playlist.id = user_with_playlist.pid").
		Where("user_with_playlist.user_id =?", userId).
		Select("user_with_playlist.pid as pid, playlist.*").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&listwith).Error

	fmt.Println(listwith)
	if len(listwith) <= 0 {
		return
	}

	for _, v := range listwith {
		elem := &model.Playlist{
			Id: v.Pid,
			Name: v.Name,
			CoverImgUrl: v.CoverImgUrl,
			CreatorUid: v.CreatorUid,
			PlayCount: v.PlayCount,
			CommentCount: v.CommentCount,
			ShareCount: v.ShareCount,
			TrackCount: v.TrackCount,
		}
		list = append(list, elem)
	}
	return
}

func (d *Dao) GetPlaylistDetailInfo(ctx context.Context, id int64) (res model.Playlist, err error) {
	if err = d.rsDBReader.Model(&model.Playlist{}).Where("id=?", id).Select("*").Find(&res).Error; err != nil {
		return
	}

	var playlistWithSong []model.PlaylistWithSong
	d.rsDBReader.Model(&model.PlaylistWithSong{}).
		Where("pid=?", id).
		Select("*").
		Find(&playlistWithSong)

	for _, item := range playlistWithSong {
		res.Sids = append(res.Sids, item.Sid)
	}

	fmt.Println("表查询结果")
	fmt.Println(res)

	return
}