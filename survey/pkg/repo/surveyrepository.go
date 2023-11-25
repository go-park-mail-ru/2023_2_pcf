package repo

import (
	"AdHub/pkg/db"
	"AdHub/survey/pkg/entities"
	"database/sql"
	"text/template"
)

type SurveyRepository struct {
	store db.DbInterface
}

func NewSurveyRepo(DB db.DbInterface) (*SurveyRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}

	r := &SurveyRepository{
		store: st,
	}

	return r, nil
}

func (r *SurveyRepository) Create(s *entities.Survey) (*entities.Survey, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"survey\" (question, type) VALUES($1, $2) RETURNING id;",
		s.Question, s.Type,
	).Scan(&s.Id); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *SurveyRepository) Remove(id int) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"survey\" WHERE id=$1;", id); err != nil {
		return err
	}

	return nil
}

func (r *SurveyRepository) Update(s *entities.Survey) error {
	_, err := r.store.Db().Exec(
		"UPDATE \"survey\" SET type=$1, question=$2 WHERE id=$3;",
		s.Type, s.Question, s.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *SurveyRepository) getOne(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, type, question FROM \"survey\" WHERE id=$1;", id)
}

func (r *SurveyRepository) Read(id int) (*entities.Survey, error) {
	rows, err := r.getOne(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var surv *entities.Survey

	for rows.Next() {
		s := &entities.Survey{}
		err = rows.Scan(&s.Id, &s.Type, &s.Question)
		if err != nil {
			return nil, err
		}

		s.Question = template.HTMLEscapeString(s.Question)

		surv = s
	}

	return surv, nil
}
