package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/db"
	"database/sql"
	"fmt"
)

type BalanceRepository struct {
	store db.DbInterface
}

func NewBalanceRepoMock(DB db.DbInterface) (*BalanceRepository, error) {
	r := &BalanceRepository{
		store: DB,
	}

	return r, nil
}

func NewBalanceRepo(DB db.DbInterface) (*BalanceRepository, error) {
	st, err := DB.Open()
	if err != nil {
		return nil, err
	}

	r := &BalanceRepository{
		store: st,
	}

	return r, nil
}

func (r *BalanceRepository) Create(balance *entities.Balance) (*entities.Balance, error) {
	if err := r.store.Db().QueryRow(
		"INSERT INTO \"balance\" (total_balance, reserved_balance, available_balance) VALUES($1, $2, $3) RETURNING id;",
		balance.Total_balance, balance.Reserved_balance, balance.Available_balance,
	).Scan(&balance.Id); err != nil {
		return nil, err
	}

	return balance, nil
}

func (r *BalanceRepository) Remove(id int) error {
	if _, err := r.store.Db().Exec("DELETE FROM \"balance\" WHERE id=$1;", id); err != nil {
		return err
	}

	return nil
}

func (r *BalanceRepository) get(id int) (*sql.Rows, error) {
	return r.store.Db().Query("SELECT id, total_balance, reserved_balance, available_balance FROM \"balance\" WHERE id=$1;", id)
}

func (r *BalanceRepository) Update(balance *entities.Balance) error {
	_, err := r.store.Db().Exec(
		"UPDATE \"balance\" SET total_balance=$1, reserved_balance=$2, available_balance=$3 WHERE id=$4;",
		balance.Total_balance, balance.Reserved_balance, balance.Available_balance, balance.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *BalanceRepository) Read(id int) (*entities.Balance, error) {
	rows, err := r.get(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	balance := &entities.Balance{}

	for rows.Next() {
		err := rows.Scan(&balance.Id, &balance.Total_balance, &balance.Reserved_balance, &balance.Available_balance)
		if err != nil {
			return nil, err
		}
	}
	if balance.Id == 0 {
		return nil, fmt.Errorf("Balance not found")
	}
	return balance, nil
}
