package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	store db.DbInterface
}

func NewUserRepoMock(DB db.DbInterface) (*UserRepository, error) {
	r := &UserRepository{
		store: DB,
	}
	return r, nil
}

func NewUserRepo(DB db.DbInterface) (*UserRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}
	r := &UserRepository{
		store: st,
	}
	return r, nil
}

func (r *UserRepository) Create(user *entities.User) (*entities.User, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"user\" (login, password, f_name, l_name, s_name, balance_id) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;",
		user.Login, user.Password, user.FName, user.LName, user.SName, user.BalanceId,
	).Scan(&user.Id); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Remove(login string) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"user\" WHERE login=$1;", login); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) get(login string) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, login, password, f_name, l_name, s_name, balance_id FROM \"user\" WHERE login=$1;", login)
}

func (r *UserRepository) Update(user *entities.User) error {
	_, err := r.store.Db().Exec(
		"UPDATE \"user\" SET login=$1, password=$2, f_name=$3, l_name=$4, s_name=$5, balance_id=$6 WHERE id=$7;",
		user.Login, user.Password, user.FName, user.LName, user.SName, user.BalanceId, user.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Read(login string) (*entities.User, error) {
	rows, err := r.get(login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user := &entities.User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Login, &user.Password, &user.FName, &user.LName, &user.SName, &user.BalanceId)
		if err != nil {
			return nil, err
		}
	}
	if user.Id == 0 {
		return nil, fmt.Errorf("User not found")
	}
	return user, nil
}
