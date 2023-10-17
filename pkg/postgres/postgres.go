package postgres

import (
	"AdHub/internal/interfaces"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Pg struct {
	db *sql.DB
}

func New() *Pg {
	return &Pg{}
}

func (s *Pg) Db() *sql.DB {
	return s.db
}

func (s *Pg) Open() (interfaces.Db, error) {
	db, err := sql.Open("postgres", "user=postgres password=zxc123 host=db port=5432 dbname=adhub sslmode=disable")
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
