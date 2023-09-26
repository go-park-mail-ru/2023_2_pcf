package store

import (
	"AdHub/internal/app/models"
	"database/sql"
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

func (r *UserRepository) Remove(mail string) error {
	if _, err := r.store.db.Exec("DELETE FROM user WHERE login=$1", mail); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Get(mail string) (*models.User, error) {
	rows, err := r.store.db.Query("SELECT * FROM user WHERE login=$1 RETURNING id, password", mail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	var password string

	for rows.Next() {
		if err := rows.Scan(&id, &password); err != nil {
			return nil, err
		}

		return &models.User{
			Id:       id,
			Login:    mail,
			Password: password,
		}, nil
	}

	return nil, sql.ErrNoRows
}
