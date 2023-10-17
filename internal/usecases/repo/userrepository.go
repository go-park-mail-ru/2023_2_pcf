package repo

import (
	"AdHub/internal/entities"
	"AdHub/internal/interfaces"
	"database/sql"
	"fmt"
	"log"
)

type UserRepository struct {
	store interfaces.Db
}

func NewUserRepo(DB interfaces.Db) (*UserRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}

	r := &UserRepository{
		store: st,
	}

	return r, nil
}

func (r *UserRepository) Create(s *entities.User) (*entities.User, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"user\" (login, password, f_name, l_name) VALUES($1, $2, $3, $4) RETURNING id;",
		s.Login, s.Password, s.FName, s.LName,
	).Scan(&s.Id); err != nil {
		log.Panic(err)
		return nil, err
	}

	return s, nil
}

func (r *UserRepository) Remove(mail string) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"user\" WHERE login=$1;", mail); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Get(mail string) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, login, password, f_name, l_name FROM \"user\" WHERE login=$1;", mail)
}

func (r *UserRepository) Update(s *entities.User) error {
	_, err := r.store.Db().Exec(
		"UPDATE \"user\" SET login=$1, password=$2, f_name=$3, l_name=$4 WHERE id=$5;",
		s.Login, s.Password, s.FName, s.LName, s.Id,
	)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}

func (r *UserRepository) Read(mail string) (*entities.User, error) {
	rows, err := r.Get(mail)
	if err != nil {
		log.Printf("Error GET user")
		return nil, err
	}

	defer rows.Close()

	user := &entities.User{} // Initialize user

	for rows.Next() {
		// Assign values to user
		err := rows.Scan(&user.Id, &user.Login, &user.Password, &user.FName, &user.LName)
		if err != nil {
			log.Printf("Error scan rows User")
			return nil, err
		}
	}
	if len(user.Login) == 0 {
		return nil, fmt.Errorf("User not found")
	}
	return user, nil
}
