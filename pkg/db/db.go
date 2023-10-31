package db

import "database/sql"

//go:generate /Users/bincom/go/bin/mockgen -source=db.go -destination=mock_db/mock.go
type DbInterface interface {
	Db() *sql.DB
	Open() (DbInterface, error)
	Close()
}
