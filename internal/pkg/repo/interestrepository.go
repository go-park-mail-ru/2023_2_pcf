package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"fmt"
	"log"
)

type InterestRepository struct {
	store db.DbInterface
}

func NewInterestRepoMock(DB db.DbInterface) (*InterestRepository, error) {
	r := &InterestRepository{
		store: DB,
	}

	return r, nil
}

func NewInterestRepo(DB db.DbInterface) (*InterestRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}

	r := &InterestRepository{
		store: st,
	}

	return r, nil
}

func (r *InterestRepository) Create(interest *entities.Interest) (*entities.Interest, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"interest\" (name) VALUES($1) RETURNING id;",
		interest.Name,
	).Scan(&interest.Id); err != nil {
		return nil, err
	}

	return interest, nil
}

func (r *InterestRepository) Remove(id int) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"interest\" WHERE id=$1;", id); err != nil {
		return err
	}

	return nil
}

func (r *InterestRepository) get(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, name FROM \"interest\" WHERE id=$1;", id)
}

func (r *InterestRepository) Update(interest *entities.Interest) error {
	_, err := r.store.Db().Exec(
		"UPDATE \"interest\" SET name=$1 WHERE id=$2;",
		interest.Name, interest.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *InterestRepository) Read(id int) (*entities.Interest, error) {
	rows, err := r.get(id)
	if err != nil {
		log.Printf("Error GET interest")
		return nil, err
	}

	defer rows.Close()

	interest := &entities.Interest{}

	for rows.Next() {
		err := rows.Scan(&interest.Id, &interest.Name)
		if err != nil {
			return nil, err
		}
	}
	if interest.Id == 0 {
		return nil, fmt.Errorf("Interest not found")
	}
	return interest, nil
}
