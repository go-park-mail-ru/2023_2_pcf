package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"fmt"
	"log"
)

type TagRepository struct {
	store db.DbInterface
}

func NewTagRepoMock(DB db.DbInterface) (*TagRepository, error) {
	r := &TagRepository{
		store: DB,
	}
	return r, nil
}

func NewTagRepo(DB db.DbInterface) (*TagRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}

	r := &TagRepository{
		store: st,
	}
	return r, nil
}

func (r *TagRepository) Create(tag *entities.Tag) (*entities.Tag, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"tags\" (name) VALUES($1) RETURNING id;",
		tag.Name,
	).Scan(&tag.Id); err != nil {
		return nil, err
	}
	return tag, nil
}

func (r *TagRepository) Remove(id int) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"tags\" WHERE id=$1;", id); err != nil {
		return err
	}
	return nil
}

func (r *TagRepository) get(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, name FROM \"tags\" WHERE id=$1;", id)
}

func (r *TagRepository) Update(tag *entities.Tag) error {
	_, err := r.store.Db().Exec(
		"UPDATE \"tags\" SET name=$1 WHERE id=$2;",
		tag.Name, tag.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *TagRepository) Read(id int) (*entities.Tag, error) {
	rows, err := r.get(id)
	if err != nil {
		log.Printf("Error GET tag")
		return nil, err
	}

	defer rows.Close()

	tag := &entities.Tag{}

	for rows.Next() {
		err := rows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return nil, err
		}
	}
	if tag.Id == 0 {
		return nil, fmt.Errorf("Tag not found")
	}
	return tag, nil
}
