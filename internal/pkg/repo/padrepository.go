package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"text/template"
)

type PadRepository struct {
	store db.DbInterface
}

func NewPadRepoMock(DB db.DbInterface) (*PadRepository, error) {
	r := &PadRepository{
		store: DB,
	}

	return r, nil
}

func NewPadRepo(DB db.DbInterface) (*PadRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}

	r := &PadRepository{
		store: st,
	}

	return r, nil
}

func (r *PadRepository) Create(pad *entities.Pad) (*entities.Pad, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"pad\" (name, description, website_link, price, target_id, owner_id, clicks, balance) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;",
		pad.Name, pad.Description, pad.Website_link, pad.Price, pad.Target_id, pad.Owner_id, pad.Clicks, pad.Balance,
	).Scan(&pad.Id); err != nil {
		return nil, err
	}

	return pad, nil
}

func (r *PadRepository) Remove(id int) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"pad\" WHERE id=$1;", id); err != nil {
		return err
	}

	return nil
}

func (r *PadRepository) Update(pad *entities.Pad) error {
	_, err := r.store.Db().Exec(
		"UPDATE \"pad\" SET name=$1, description=$2, website_link=$3, price=$4, target_id=$5, owner_id=$6, clicks=$7, balnce=$8 WHERE id=$9;",
		pad.Name, pad.Description, pad.Website_link, pad.Price, pad.Target_id, pad.Owner_id, pad.Clicks, pad.Id, pad.Balance,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PadRepository) Read(id int) ([]*entities.Pad, error) {
	rows, err := r.getList(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pads []*entities.Pad

	for rows.Next() {
		pad := &entities.Pad{}
		err := rows.Scan(&pad.Id, &pad.Clicks, &pad.Name, &pad.Description, &pad.Website_link, &pad.Price, &pad.Target_id, &pad.Owner_id, &pad.Balance)
		if err != nil {
			return nil, err
		}

		pad.Name = template.HTMLEscapeString(pad.Name)
		pad.Description = template.HTMLEscapeString(pad.Description)
		pad.Website_link = template.HTMLEscapeString(pad.Website_link)

		pads = append(pads, pad)
	}

	return pads, nil
}

func (r *PadRepository) getList(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, clicks, name, description, website_link, price, target_id, owner_id, balance FROM \"pad\" WHERE owner_id=$1;", id)
}

func (r *PadRepository) getOne(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, clicks, name, description, website_link, price, target_id, owner_id, balance FROM \"pad\" WHERE id=$1;", id)
}

func (r *PadRepository) Get(id int) (*entities.Pad, error) {
	rows, err := r.getOne(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pads *entities.Pad

	for rows.Next() {
		pad := &entities.Pad{}
		err = rows.Scan(&pad.Id, &pad.Clicks, &pad.Name, &pad.Description, &pad.Website_link, &pad.Price, &pad.Target_id, &pad.Owner_id, &pad.Balance)
		if err != nil {
			return nil, err
		}

		pad.Name = template.HTMLEscapeString(pad.Name)
		pad.Description = template.HTMLEscapeString(pad.Description)
		pad.Website_link = template.HTMLEscapeString(pad.Website_link)

		pads = pad
	}

	return pads, nil
}
