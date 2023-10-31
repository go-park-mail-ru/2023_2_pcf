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

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewAdRepoMock(dbInterface)
	require.NoError(t, err)

	ad := &entities.Ad{
		Name:        "Test Ad",
		Description: "Test Description",
		Sector:      "Test Sector",
		Owner_id:    1,
	}

	mock.ExpectQuery("INSERT INTO \"ad\" (.+) RETURNING id;").
		WithArgs(ad.Name, ad.Description, ad.Sector, ad.Owner_id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdAd, err := repo.Create(ad)

	require.NoError(t, err)

	assert.Equal(t, 1, createdAd.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRemove(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewAdRepoMock(dbInterface)
	require.NoError(t, err)

	mock.ExpectExec(`DELETE FROM "ad" WHERE id=\$1;`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove(1)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewAdRepoMock(dbInterface)
	require.NoError(t, err)

	ad := &entities.Ad{
		Id:          1,
		Name:        "Updated Ad",
		Description: "Updated Description",
		Sector:      "Updated Sector",
		Owner_id:    1,
	}

	mock.ExpectExec(`UPDATE "ad" SET name=\$1, description=\$2, sector=\$3, owner_id=\$4 WHERE id=\$5;`).
		WithArgs(ad.Name, ad.Description, ad.Sector, ad.Owner_id, ad.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(ad)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRead(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewAdRepoMock(dbInterface)
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"name", "description", "sector", "id"}).
		AddRow("Ad1", "Description1", "Sector1", 1).
		AddRow("Ad2", "Description2", "Sector2", 2)

	mock.ExpectQuery(`SELECT name, description, sector, id FROM "ad" WHERE owner_id=\$1;`).
		WillReturnRows(rows)

	ads, err := repo.Read(1)

	require.NoError(t, err)

	assert.Len(t, ads, 2)

	assert.Equal(t, "Ad1", ads[0].Name)
	assert.Equal(t, "Description1", ads[0].Description)
	assert.Equal(t, "Sector1", ads[0].Sector)
	assert.Equal(t, 1, ads[0].Id)

	assert.Equal(t, "Ad2", ads[1].Name)
	assert.Equal(t, "Description2", ads[1].Description)
	assert.Equal(t, "Sector2", ads[1].Sector)
	assert.Equal(t, 2, ads[1].Id)

	require.NoError(t, mock.ExpectationsWereMet())
}
