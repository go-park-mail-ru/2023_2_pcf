package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Pg struct {
	connect string
	db      *sql.DB
}

func New(connect string) *Pg {
	return &Pg{connect: connect}
}

func NewMock(db *sql.DB) *Pg {
	return &Pg{
		db: db,
	}
}

func (s *Pg) Db() *sql.DB {
	return s.db
}

func (s *Pg) Open() (DbInterface, error) {
	db, err := sql.Open("postgres", s.connect)
	if err != nil {
		log.Printf("open")
		return nil, err
	}

	s.db = db

	return s, nil
}

func (s *Pg) Close() {
	s.db.Close()
}
