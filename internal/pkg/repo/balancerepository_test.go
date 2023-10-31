package repo

import (
	"AdHub/internal/pkg/entities"
	pg "AdHub/pkg/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewBalanceRepoMock(dbInterface)
	require.NoError(t, err)

	balance := &entities.Balance{
		Total_balance:     1000.0,
		Reserved_balance:  200.0,
		Available_balance: 800.0,
	}

	mock.ExpectQuery("INSERT INTO \"balance\" (.+) RETURNING id;").
		WithArgs(balance.Total_balance, balance.Reserved_balance, balance.Available_balance).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdBalance, err := repo.Create(balance)

	require.NoError(t, err)

	assert.Equal(t, 1, createdBalance.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRemoveBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewBalanceRepoMock(dbInterface)
	require.NoError(t, err)

	mock.ExpectExec(`DELETE FROM "balance" WHERE id=\$1;`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove(1)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewBalanceRepoMock(dbInterface)
	require.NoError(t, err)

	balance := &entities.Balance{
		Id:                1,
		Total_balance:     1500.0,
		Reserved_balance:  300.0,
		Available_balance: 1200.0,
	}

	mock.ExpectExec(`UPDATE "balance" SET total_balance=\$1, reserved_balance=\$2, available_balance=\$3 WHERE id=\$4;`).
		WithArgs(balance.Total_balance, balance.Reserved_balance, balance.Available_balance, balance.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(balance)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReadBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewBalanceRepoMock(dbInterface)
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "total_balance", "reserved_balance", "available_balance"}).
		AddRow(1, 1000.0, 200.0, 800.0)

	mock.ExpectQuery(`SELECT id, total_balance, reserved_balance, available_balance FROM "balance" WHERE id=\$1;`).
		WillReturnRows(rows)

	balance, err := repo.Read(1)

	require.NoError(t, err)

	assert.Equal(t, 1, balance.Id)
	assert.Equal(t, 1000.0, balance.Total_balance)
	assert.Equal(t, 200.0, balance.Reserved_balance)
	assert.Equal(t, 800.0, balance.Available_balance)

	require.NoError(t, mock.ExpectationsWereMet())
}
