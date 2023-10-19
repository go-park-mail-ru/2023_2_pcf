package db

import "database/sql"

type DbInterface interface {
	Db() *sql.DB
	Open() (DbInterface, error)
	Close()
}
