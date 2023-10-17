package pg

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Pg struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Pg {
	return &Pg{
		config: config,
	}
}

func (s *Pg) Open() (*Pg, error) {
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
