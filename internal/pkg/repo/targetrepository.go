package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"fmt"
	"log"
	"strings"
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
	tagsStr := strings.Join(target.Tags, ", ")
	regionsStr := strings.Join(target.Regions, ", ")
	interestsStr := strings.Join(target.Interests, ", ")
	keysStr := strings.Join(target.Keys, ", ")

	if err := r.store.Db().QueryRow(
		"INSERT INTO \"target\" (name, owner_id, gender, min_age, max_age, tags, regions, interests, keys) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;",
		target.Name, target.Owner_id, target.Gender, target.Min_age, target.Max_age, tagsStr, regionsStr, interestsStr, keysStr,
	).Scan(&target.Id); err != nil {
		return nil, err
	}

	return target, nil
}

func (r *TargetRepository) Update(target *entities.Target) error {
	tagsStr := strings.Join(target.Tags, ", ")
	regionsStr := strings.Join(target.Regions, ", ")
	interestsStr := strings.Join(target.Interests, ", ")
	keysStr := strings.Join(target.Keys, ", ")

	_, err := r.store.Db().Exec(
		"UPDATE \"target\" SET name=$1, owner_id=$2, gender=$3, min_age=$4, max_age=$5, tags=$6, regions=$7, interests=$8, keys=$9 WHERE id=$10;",
		target.Name, target.Owner_id, target.Gender, target.Min_age, target.Max_age, tagsStr, regionsStr, interestsStr, keysStr, target.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *TargetRepository) Remove(id int) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"targets\" WHERE id=$1;", id); err != nil {
		return err
	}

	return nil
}

func (r *TargetRepository) get(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, name, owner_id, gender, min_age, max_age, tags, regions, interests, keys FROM \"target\" WHERE id=$1;", id)
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
		var tagsStr, regionsStr, interestsStr, keysStr string
		err := rows.Scan(&target.Id, &target.Name, &target.Owner_id, &target.Gender, &target.Min_age, &target.Max_age, &tagsStr, &regionsStr, &interestsStr, &keysStr)
		if err != nil {
			return nil, err
		}

		target.Tags = splitTags(tagsStr)
		target.Regions = splitTags(regionsStr)
		target.Interests = splitTags(interestsStr)
		target.Keys = splitTags(keysStr)
	}

	if target.Id == 0 {
		return nil, fmt.Errorf("Target not found")
	}

	return target, nil
}

func splitTags(input string) []string {
	tags := strings.Split(input, ",")
	for i := range tags {
		tags[i] = strings.TrimSpace(tags[i])
	}
	return tags
}
