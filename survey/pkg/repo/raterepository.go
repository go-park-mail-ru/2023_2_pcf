package repo

import (
	"AdHub/pkg/db"
	"AdHub/survey/pkg/entities"
	"database/sql"
)

type RateRepository struct {
	store db.DbInterface
}

func NewRateRepo(DB db.DbInterface) (*RateRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}

	r := &RateRepository{
		store: st,
	}

	return r, nil
}

func (r *RateRepository) Create(s *entities.Rate) (*entities.Rate, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"rate\" (user_id, rate, survey_id) VALUES($1, $2, $3) RETURNING id;",
		s.User_id, s.Rate, s.Survey_id,
	).Scan(&s.Id); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *RateRepository) Read(id int) ([]*entities.Rate, error) {
	rows, err := r.getList(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rates []*entities.Rate

	for rows.Next() {
		rate := &entities.Rate{}
		err := rows.Scan(&rate.Id, &rate.User_id, &rate.Rate, &rate.Survey_id)
		if err != nil {
			return nil, err
		}

		rates = append(rates, rate)
	}

	return rates, nil
}

func (r *RateRepository) getList(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, user_id, rate, survey_id FROM \"rate\" WHERE survey_id=$1;", id)
}
