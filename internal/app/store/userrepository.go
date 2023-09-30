package store

import (
	"AdHub/internal/app/models"
	"database/sql"
	"log"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(s *models.User) (*models.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO \"user\" (login, password) VALUES($1, $2) RETURNING id;",
		s.Login, s.Password,
	).Scan(&s.Id); err != nil {
		log.Panic(err)
		return nil, err
	}

	return s, nil
}

func (r *UserRepository) Remove(mail string) error {
	if _, err := r.store.db.Exec("DELETE FROM \"user\" WHERE login=$1;", mail); err != nil {
		return err
	}

	return nil
}

/*func (r *UserRepository) Get(mail string) (models.User, error) {
	rows, err := r.store.db.Query("SELECT * FROM user WHERE login=$1 RETURNING id, password", mail)
	if err != nil {
		//error
	}
	defer rows.Close()

	var id int
	var login, password string

	if err := rows.Scan(&id, &login, &password); err != nil {
		//
	}

	return models.User{
		Id:       id,
		Login:    login,
		Password: password,
	}, nil
}*/

func (r *UserRepository) Get(mail string) (*sql.Rows, error) {
	return r.store.db.Query("SELECT id, login, password FROM \"user\" WHERE login=$1;", mail)
}

func (r *UserRepository) Update(s *models.User) error {
	_, err := r.store.db.Exec(
		"UPDATE \"user\" SET login=$1, password=$2 WHERE id=$3;",
		s.Login, s.Password, s.Id,
	)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
