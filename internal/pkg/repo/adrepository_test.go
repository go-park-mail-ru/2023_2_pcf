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

func TestAdRepository_Create(t *testing.T) {
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
		Name:         "Test Ad",
		Description:  "Test Description",
		Website_link: "https://example.com",
		Budget:       100.0,
		Target_id:    1,
		Image_link:   "https://example.com/image.jpg",
		Owner_id:     1,
		Click_cost:   0.5,
	}

	mock.ExpectQuery(`INSERT INTO "ad" (.+) RETURNING id;`).
		WithArgs(ad.Name, ad.Description, ad.Website_link, ad.Budget, ad.Target_id, ad.Image_link, ad.Owner_id, ad.Click_cost).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdAd, err := repo.Create(ad)

	require.NoError(t, err)

	assert.Equal(t, 1, createdAd.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAdRepository_Remove(t *testing.T) {
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
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove(1)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAdRepository_Read(t *testing.T) {
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

	rows := sqlmock.NewRows([]string{"id", "name", "description", "website_link", "budget", "target_id", "image_link", "owner_id", "click_cost"}).
		AddRow(1, "Ad1", "Description1", "https://example1.com", 100.0, 1, "https://example1.com/image.jpg", 1, 0.5).
		AddRow(2, "Ad2", "Description2", "https://example2.com", 200.0, 2, "https://example2.com/image.jpg", 2, 0.8)

	mock.ExpectQuery(`SELECT id, name, description, website_link, budget, target_id, image_link, owner_id, click_cost FROM "ad" WHERE owner_id=\$1;`).
		WithArgs(1).
		WillReturnRows(rows)

	ads, err := repo.Read(1)

	require.NoError(t, err)

	assert.Len(t, ads, 2)

	assert.Equal(t, "Ad1", ads[0].Name)
	assert.Equal(t, "Description1", ads[0].Description)
	assert.Equal(t, "https://example1.com", ads[0].Website_link)
	assert.Equal(t, 100.0, ads[0].Budget)
	assert.Equal(t, 1, ads[0].Target_id)
	assert.Equal(t, "https://example1.com/image.jpg", ads[0].Image_link)
	assert.Equal(t, 1, ads[0].Owner_id)
	assert.Equal(t, 0.5, ads[0].Click_cost)
	assert.Equal(t, 1, ads[0].Id)

	assert.Equal(t, "Ad2", ads[1].Name)
	assert.Equal(t, "Description2", ads[1].Description)
	assert.Equal(t, "https://example2.com", ads[1].Website_link)
	assert.Equal(t, 200.0, ads[1].Budget)
	assert.Equal(t, 2, ads[1].Target_id)
	assert.Equal(t, "https://example2.com/image.jpg", ads[1].Image_link)
	assert.Equal(t, 2, ads[1].Owner_id)
	assert.Equal(t, 0.8, ads[1].Click_cost)
	assert.Equal(t, 2, ads[1].Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAdRepository_Get(t *testing.T) {
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

	rows := sqlmock.NewRows([]string{"id", "name", "description", "website_link", "budget", "target_id", "image_link", "owner_id", "click_cost"}).
		AddRow(1, "Ad1", "Description1", "https://example1.com", 100.0, 1, "https://example1.com/image.jpg", 1, 0.5)

	mock.ExpectQuery(`SELECT id, name, description, website_link, budget, target_id, image_link, owner_id, click_cost FROM "ad" WHERE id=\$1;`).
		WithArgs(1).
		WillReturnRows(rows)

	ad, err := repo.Get(1)

	require.NoError(t, err)
	assert.Equal(t, "Ad1", ad.Name)
	assert.Equal(t, "Description1", ad.Description)
	assert.Equal(t, "https://example1.com", ad.Website_link)
	assert.Equal(t, 100.0, ad.Budget)
	assert.Equal(t, 1, ad.Target_id)
	assert.Equal(t, "https://example1.com/image.jpg", ad.Image_link)
	assert.Equal(t, 1, ad.Owner_id)
	assert.Equal(t, 0.5, ad.Click_cost)
	assert.Equal(t, 1, ad.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}
