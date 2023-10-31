package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"fmt"
	"log"
)

type RegionRepository struct {
	store db.DbInterface
}

func NewRegionRepoMock(DB db.DbInterface) (*RegionRepository, error) {
	r := &RegionRepository{
		store: DB,
	}
	return r, nil
}

func NewRegionRepo(DB db.DbInterface) (*RegionRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}

	r := &RegionRepository{
		store: st,
	}

	return r, nil
}

func (r *RegionRepository) Create(region *entities.Region) (*entities.Region, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"regions\" (name) VALUES($1) RETURNING id;",
		region.Name,
	).Scan(&region.Id); err != nil {
		log.Panic(err)
		return nil, err
	}

	return region, nil
}

func (r *RegionRepository) Remove(id int) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"regions\" WHERE id=$1;", id); err != nil {
		return err
	}

	return nil
}

func (r *RegionRepository) get(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, name FROM \"regions\" WHERE id=$1;", id)
}

func (r *RegionRepository) Update(region *entities.Region) error {
	_, err := r.store.Db().Exec(
		"UPDATE \"regions\" SET name=$1 WHERE id=$2;",
		region.Name, region.Id,
	)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}

func (r *RegionRepository) Read(id int) (*entities.Region, error) {
	rows, err := r.get(id)
	if err != nil {
		log.Printf("Error GET region")
		return nil, err
	}

	defer rows.Close()

	region := &entities.Region{}

	for rows.Next() {
		err := rows.Scan(&region.Id, &region.Name)
		if err != nil {
			log.Printf("Error scan rows Region")
			return nil, err
		}
	}
	if region.Id == 0 {
		return nil, fmt.Errorf("Region not found")
	}
	return region, nil
}
