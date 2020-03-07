package conf

import (
	"flag"

	"github.com/longhao/music/library/database/sql"
	"github.com/longhao/music/library/database/orm"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	Conf = &Config{}
)

type Config struct {
	MySQL *sql.Config
	DB *DB
}

type DB struct {
	Resource *orm.Config
	ResourceReader *orm.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")

}

func Init() error {
	return local()
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}