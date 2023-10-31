package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"fmt"
	"log"
)

type TargetRepository struct {
	store db.DbInterface
}

func NewTargetRepoMock(DB db.DbInterface) (*TargetRepository, error) {
	r := &TargetRepository{
		store: DB,
	}

	return r, nil
}

func NewTargetRepo(DB db.DbInterface) (*TargetRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}

	r := &TargetRepository{
		store: st,
	}

	return r, nil
}

func (r *TargetRepository) Create(target *entities.Target) (*entities.Target, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"targets\" (name, owner_id, gender, min_age, max_age) VALUES($1, $2, $3, $4, $5) RETURNING id;",
		target.Name, target.Owner_id, target.Gender, target.Min_age, target.Max_age,
	).Scan(&target.Id); err != nil {
		return nil, err
	}

	return target, nil
}

func (r *TargetRepository) Remove(id int) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"targets\" WHERE id=$1;", id); err != nil {
		return err
	}

	return nil
}

func (r *TargetRepository) get(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, name, owner_id, gender, min_age, max_age FROM \"targets\" WHERE id=$1;", id)
}

func (r *TargetRepository) Update(target *entities.Target) error {
	_, err := r.store.Db().Exec(
		"UPDATE \"targets\" SET name=$1, owner_id=$2, gender=$3, min_age=$4, max_age=$5 WHERE id=$6;",
		target.Name, target.Owner_id, target.Gender, target.Min_age, target.Max_age, target.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *TargetRepository) Read(id int) (*entities.Target, error) {
	rows, err := r.get(id)
	if err != nil {
		log.Printf("Error GET target")
		return nil, err
	}

	defer rows.Close()

	target := &entities.Target{}

	for rows.Next() {
		err := rows.Scan(&target.Id, &target.Name, &target.Owner_id, &target.Gender, &target.Min_age, &target.Max_age)
		if err != nil {
			return nil, err
		}
	}
	if target.Id == 0 {
		return nil, fmt.Errorf("Target not found")
	}
	return target, nil
}

func (r *TargetRepository) GetTargetTags(targetID int) ([]entities.Tag, error) {
	rows, err := r.store.Db().Query(`
        SELECT t.id, t.name
        FROM "tags" t
        JOIN "target_tags" tt ON t.id = tt.tag_id
        WHERE tt.target_id = $1;
    `, targetID)
	if err != nil {
		log.Printf("Error getting target tags")
		return nil, err
	}
	defer rows.Close()

	var tags []entities.Tag

	for rows.Next() {
		tag := entities.Tag{}
		if err := rows.Scan(&tag.Id, &tag.Name); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *TargetRepository) GetTargetRegions(targetID int) ([]entities.Region, error) {
	rows, err := r.store.Db().Query(`
        SELECT r.id, r.name
        FROM "regions" r
        JOIN "target_regions" tr ON r.id = tr.region_id
        WHERE tr.target_id = $1;
    `, targetID)
	if err != nil {
		log.Printf("Error getting target regions")
		return nil, err
	}
	defer rows.Close()

	var regions []entities.Region

	for rows.Next() {
		region := entities.Region{}
		if err := rows.Scan(&region.Id, &region.Name); err != nil {
			return nil, err
		}
		regions = append(regions, region)
	}

	return regions, nil
}

func (r *TargetRepository) GetTargetInterests(targetID int) ([]entities.Interest, error) {
	rows, err := r.store.Db().Query(`
        SELECT i.id, i.name
        FROM "interests" i
        JOIN "target_interests" ti ON i.id = ti.interest_id
        WHERE ti.target_id = $1;
    `, targetID)
	if err != nil {
		log.Printf("Error getting target interests")
		return nil, err
	}
	defer rows.Close()

	var interests []entities.Interest

	for rows.Next() {
		interest := entities.Interest{}
		if err := rows.Scan(&interest.Id, &interest.Name); err != nil {
			return nil, err
		}
		interests = append(interests, interest)
	}

	return interests, nil
}
