package store

import (
	"AdHub/internal/app/models"
	"database/sql"
	"log"
)

type AdRepository struct {
	store *Store
}

func (r *AdRepository) Create(s *models.Ad) (*models.Ad, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO \"ad\" (name, description, sector, owner_id) VALUES($1, $2, $3, $4) RETURNING id;",
		s.Name, s.Description, s.Sector, s.Owner_id,
	).Scan(&s.Id); err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}

	return s, nil
}

func (r *AdRepository) Remove(id int) error {
	if _, err := r.store.db.Exec("DELETE FROM \"ad\" WHERE id=$1;", id); err != nil {
		return err
	}

	return nil
}

func (r *AdRepository) Get(id int) (*sql.Rows, error) {
	return r.store.db.Query("SELECT name, description, sector, owner_id FROM \"ad\" WHERE id=$1;", id)
}

func (r *AdRepository) GetList(id int) (*sql.Rows, error) {
	return r.store.db.Query("SELECT name, description, sector, id FROM \"ad\" WHERE owner_id=$1;", id)
}

func (r *AdRepository) Update(s *models.Ad) error {
	_, err := r.store.db.Exec(
		"UPDATE \"ad\" SET name=$1, description=$2, sector=$3, owner_id=$4 WHERE id=$5;",
		s.Name, s.Description, s.Sector, s.Owner_id, s.Id,
	)
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}

	return nil
}

func (r *AdRepository) Read(id int) ([]*models.Ad, error) {
	rows, err := r.GetList(id)
	if err != nil {
		log.Printf("Error GET Ads")
		return nil, err
	}

	defer rows.Close()

	var ads []*models.Ad

	for rows.Next() {
		ad := &models.Ad{}
		err := rows.Scan(&ad.Name, &ad.Description, &ad.Sector, &ad.Id)
		ad.Owner_id = id
		if err != nil {
			log.Printf("Error scan rows Ads")
			return nil, err
		}
		ads = append(ads, ad)
	}

	return ads, nil
}
