package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
	adRepository   *AdRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", "user=postgres password=zxc123 host=db port=5432 dbname=adhub sslmode=disable")
	if err != nil {
		log.Printf("open")
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

// Чтобы обратиться допустим s.User().Add(us User)
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Ad() *AdRepository {
	if s.adRepository != nil {
		return s.adRepository
	}

	s.adRepository = &AdRepository{
		store: s,
	}

	return s.adRepository
}
