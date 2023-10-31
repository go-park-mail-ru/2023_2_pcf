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

func TestCreateInterest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewInterestRepoMock(dbInterface)
	require.NoError(t, err)

	interest := &entities.Interest{
		Name: "Test Interest",
	}

	mock.ExpectQuery("INSERT INTO \"interest\" (.+) RETURNING id;").
		WithArgs(interest.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdInterest, err := repo.Create(interest)

	require.NoError(t, err)

	assert.Equal(t, 1, createdInterest.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRemoveInterest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewInterestRepoMock(dbInterface)
	require.NoError(t, err)

	mock.ExpectExec(`DELETE FROM "interest" WHERE id=\$1;`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove(1)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateInterest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewInterestRepoMock(dbInterface)
	require.NoError(t, err)

	interest := &entities.Interest{
		Id:   1,
		Name: "Updated Interest",
	}

	mock.ExpectExec(`UPDATE "interest" SET name=\$1 WHERE id=\$2;`).
		WithArgs(interest.Name, interest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(interest)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReadInterest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewInterestRepoMock(dbInterface)
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Test Interest")

	mock.ExpectQuery(`SELECT id, name FROM "interest" WHERE id=\$1;`).
		WillReturnRows(rows)

	interest, err := repo.Read(1)

	require.NoError(t, err)

	assert.Equal(t, 1, interest.Id)
	assert.Equal(t, "Test Interest", interest.Name)

	require.NoError(t, mock.ExpectationsWereMet())
}
