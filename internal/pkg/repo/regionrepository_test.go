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

func TestCreateRegion(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewRegionRepoMock(dbInterface)
	require.NoError(t, err)

	region := &entities.Region{
		Name: "Test Region",
	}

	mock.ExpectQuery("INSERT INTO \"regions\" (.+) RETURNING id;").
		WithArgs(region.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdRegion, err := repo.Create(region)

	require.NoError(t, err)

	assert.Equal(t, 1, createdRegion.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRemoveRegion(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewRegionRepoMock(dbInterface)
	require.NoError(t, err)

	mock.ExpectExec(`DELETE FROM "regions" WHERE id=\$1;`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove(1)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateRegion(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewRegionRepoMock(dbInterface)
	require.NoError(t, err)

	region := &entities.Region{
		Id:   1,
		Name: "Updated Region",
	}

	mock.ExpectExec(`UPDATE "regions" SET name=\$1 WHERE id=\$2;`).
		WithArgs(region.Name, region.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(region)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReadRegion(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewRegionRepoMock(dbInterface)
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Test Region")

	mock.ExpectQuery(`SELECT id, name FROM "regions" WHERE id=\$1;`).
		WillReturnRows(rows)

	region, err := repo.Read(1)

	require.NoError(t, err)

	assert.Equal(t, 1, region.Id)
	assert.Equal(t, "Test Region", region.Name)

	require.NoError(t, mock.ExpectationsWereMet())
}
