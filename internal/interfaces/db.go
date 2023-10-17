package interfaces

import "database/sql"

type Db interface {
	Db() *sql.DB
	Open() (Db, error)
	Close()
}
