package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"text/template"
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
		"INSERT INTO \"ad\" (name, description, website_link, budget, target_id, image_link, owner_id, click_cost) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;",
		ad.Name, ad.Description, ad.Website_link, ad.Budget, ad.Target_id, ad.Image_link, ad.Owner_id, ad.Click_cost,
	).Scan(&ad.Id); err != nil {
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
		"UPDATE \"ad\" SET name=$1, description=$2, website_link=$3, budget=$4, target_id=$5, image_link=$6, owner_id=$7, click_cost=$8 WHERE id=$9;",
		ad.Name, ad.Description, ad.Website_link, ad.Budget, ad.Target_id, ad.Image_link, ad.Owner_id, ad.Id, ad.Click_cost,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *AdRepository) Read(id int) ([]*entities.Ad, error) {
	rows, err := r.getList(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ads []*entities.Ad

	for rows.Next() {
		ad := &entities.Ad{}
		err := rows.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Website_link, &ad.Budget, &ad.Target_id, &ad.Image_link, &ad.Owner_id, &ad.Click_cost)
		if err != nil {
			return nil, err
		}

		ad.Name = template.HTMLEscapeString(ad.Name)
		ad.Description = template.HTMLEscapeString(ad.Description)
		ad.Website_link = template.HTMLEscapeString(ad.Website_link)
		ad.Image_link = template.HTMLEscapeString(ad.Image_link)

		ads = append(ads, ad)
	}

	return ads, nil
}

func (r *AdRepository) getList(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, name, description, website_link, budget, target_id, image_link, owner_id, click_cost FROM \"ad\" WHERE owner_id=$1;", id)
}

func (r *AdRepository) getOne(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, name, description, website_link, budget, target_id, image_link, owner_id, click_cost FROM \"ad\" WHERE id=$1;", id)
}

func (r *AdRepository) getTarget(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, name, description, website_link, budget, target_id, image_link, owner_id, click_cost FROM \"ad\" WHERE target_id=$1;", id)
}

func (r *AdRepository) Get(id int) (*entities.Ad, error) {
	rows, err := r.getOne(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ads *entities.Ad

	for rows.Next() {
		ad := &entities.Ad{}
		err = rows.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Website_link, &ad.Budget, &ad.Target_id, &ad.Image_link, &ad.Owner_id, &ad.Click_cost)
		if err != nil {
			return nil, err
		}

		ad.Name = template.HTMLEscapeString(ad.Name)
		ad.Description = template.HTMLEscapeString(ad.Description)
		ad.Website_link = template.HTMLEscapeString(ad.Website_link)
		ad.Image_link = template.HTMLEscapeString(ad.Image_link)

		ads = ad
	}

	return ads, nil
}

func (r *AdRepository) ReaByTarget(id int) ([]*entities.Ad, error) {
	rows, err := r.getTarget(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ads []*entities.Ad

	for rows.Next() {
		ad := &entities.Ad{}
		err := rows.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Website_link, &ad.Budget, &ad.Target_id, &ad.Image_link, &ad.Owner_id, &ad.Click_cost)
		if err != nil {
			return nil, err
		}

		ad.Name = template.HTMLEscapeString(ad.Name)
		ad.Description = template.HTMLEscapeString(ad.Description)
		ad.Website_link = template.HTMLEscapeString(ad.Website_link)
		ad.Image_link = template.HTMLEscapeString(ad.Image_link)

		ads = append(ads, ad)
	}

	return ads, nil
}
