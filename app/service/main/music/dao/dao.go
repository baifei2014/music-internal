package dao

import (
	"database/sql"

	lsql "github.com/longhao/music/library/database/sql"
	"github.com/longhao/music/library/database/orm"
	"github.com/longhao/music/app/service/main/music/conf"

	"github.com/jinzhu/gorm"
)
type Dao struct {
	db *sql.DB
	rsDB       *gorm.DB
	rsDBReader *gorm.DB
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		db: lsql.NewMySQL(c.MySQL),
		rsDB: orm.NewMySQL(c.DB.Resource),
		rsDBReader: orm.NewMySQL(c.DB.ResourceReader),
	}
	return
}

