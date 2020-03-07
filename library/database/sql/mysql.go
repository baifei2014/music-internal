package sql

import (
	// "time"
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Config mysql config.
type Config struct {
	Addr         string          // for trace
	DSN          string          // write data source name.
	ReadDSN      []string        // read data source name.
	Active       int             // pool
	Idle         int             // pool
	// IdleTimeout  time.Duration   // connect max life time.
	// QueryTimeout time.Duration   // query sql timeout
	// ExecTimeout  time.Duration   // execute sql timeout
	// TranTimeout  time.Duration   // transaction sql timeout
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *sql.DB) {
	fmt.Println(c)
	db, err := Open(c)
	if err != nil {
		panic(err)
	}
	return
}