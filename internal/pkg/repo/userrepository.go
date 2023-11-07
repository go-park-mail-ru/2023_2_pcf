package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"fmt"
	"text/template"
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
		`INSERT INTO "user" (login, password, f_name, l_name) VALUES($1, $2, $3, $4) RETURNING id;`,
		user.Login, user.Password, user.FName, user.LName,
	).Scan(&user.Id); err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) Remove(login string) error {
	if _, err := r.store.Db().Exec(`DELETE FROM user WHERE login=$1;`, login); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) getByLogin(login string) (*sql.Rows, error) {
	return r.store.Db().Query(`SELECT id, login, password FROM "user" WHERE login=$1;`, login)
}

func (r *UserRepository) getById(id int) (*sql.Rows, error) {
	return r.store.Db().Query(`SELECT id, login, password, f_name, l_name, s_name, balance_id, avatar FROM "user" WHERE id=$1;`, id)
}

func (r *UserRepository) Update(user *entities.User) error {
	_, err := r.store.Db().Exec(
		`UPDATE "user" SET login=$1, password=$2, f_name=$3, l_name=$4, s_name=$5, avatar=$6, balance_id=$7 WHERE id=$8;`,
		user.Login, user.Password, user.FName, user.LName, user.CompanyName, user.Avatar, user.BalanceId, user.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) ReadByLogin(login string) (*entities.User, error) {
	rows, err := r.getByLogin(login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user := &entities.User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Login, &user.Password)
		if err != nil {
			return nil, err
		}
	}
	user.Login = template.HTMLEscapeString(user.Login)
	user.Password = template.HTMLEscapeString(user.Password)
	user.FName = template.HTMLEscapeString(user.FName)
	user.LName = template.HTMLEscapeString(user.LName)
	user.CompanyName = template.HTMLEscapeString(user.CompanyName)
	user.Avatar = template.HTMLEscapeString(user.Avatar)

	if user.Id == 0 {
		return nil, fmt.Errorf("User not found")
	}
	return user, nil
}

func (r *UserRepository) ReadById(id int) (*entities.User, error) {
	rows, err := r.getById(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user := &entities.User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Login, &user.Password, &user.FName, &user.LName, &user.CompanyName, &user.Avatar, &user.BalanceId)
		if err != nil {
			return nil, err
		}
	}
	user.Login = template.HTMLEscapeString(user.Login)
	user.Password = template.HTMLEscapeString(user.Password)
	user.FName = template.HTMLEscapeString(user.FName)
	user.LName = template.HTMLEscapeString(user.LName)
	user.CompanyName = template.HTMLEscapeString(user.CompanyName)
	user.Avatar = template.HTMLEscapeString(user.Avatar)

	if user.Id == 0 {
		return nil, fmt.Errorf("User not found")
	}
	return user, nil
}
