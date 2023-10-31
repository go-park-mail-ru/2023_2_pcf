package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"log"
)

type AdRepository struct {
	store db.DbInterface
}

func NewAdRepoMock(DB db.DbInterface) (*AdRepository, error) {
	r := &AdRepository{
		store: DB,
	}

	return r, nil
}

func NewAdRepo(DB db.DbInterface) (*AdRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}

	r := &AdRepository{
		store: st,
	}

	return r, nil
}

func (r *AdRepository) Create(ad *entities.Ad) (*entities.Ad, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"ad\" (name, description, website_link, budget, target_id, image_link, owner_id) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;",
		ad.Name, ad.Description, ad.Website_link, ad.Budget, ad.Target_id, ad.Image_link, ad.Owner_id,
	).Scan(&ad.Id); err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}

	return ad, nil
}

func (r *AdRepository) Remove(id int) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"ad\" WHERE id=$1;", id); err != nil {
		return err
	}

	return nil
}

func (r *AdRepository) Update(ad *entities.Ad) error {
	_, err := r.store.Db().Exec(
		"UPDATE \"ad\" SET name=$1, description=$2, website_link=$3, budget=$4, target_id=$5, image_link=$6, owner_id=$7 WHERE id=$8;",
		ad.Name, ad.Description, ad.Website_link, ad.Budget, ad.Target_id, ad.Image_link, ad.Owner_id, ad.Id,
	)
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}

	return nil
}

func (r *AdRepository) Read(id int) ([]*entities.Ad, error) {
	rows, err := r.getList(id)
	if err != nil {
		log.Printf("Error GET Ads")
		return nil, err
	}

	defer rows.Close()

	var ads []*entities.Ad

	for rows.Next() {
		ad := &entities.Ad{}
		err := rows.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Website_link, &ad.Budget, &ad.Target_id, &ad.Image_link, &ad.Owner_id)
		if err != nil {
			log.Printf("Error scan rows Ads")
			return nil, err
		}
		ads = append(ads, ad)
	}

	return ads, nil
}

func (r *AdRepository) getList(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, name, description, website_link, budget, target_id, image_link, owner_id FROM \"ad\" WHERE owner_id=$1;", id)
}
