package sql

import (
	// "time"
	"database/sql"
)

func Open(c *Config) (*sql.DB, error) {
	db, err := connect(c, c.DSN)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connect(c *Config, dataSourceName string) (*sql.DB, error) {
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(c.Active)
	d.SetMaxIdleConns(c.Idle)
	// d.SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	return d, nil
}