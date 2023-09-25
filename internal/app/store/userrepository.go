package store

import (
	"AdHub/internal/app/models"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Add(s *models.User) (*models.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO user (login, password) VALUES($1, $2) RETURNING id",
		s.Login, s.Password,
	).Scan(&s.Id); err != nil {
		return nil, err
	}

	return s, nil
}
