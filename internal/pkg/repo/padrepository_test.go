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

func TestPadRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewPadRepoMock(dbInterface)
	require.NoError(t, err)

	pad := &entities.Pad{
		Name:         "Test Pad",
		Description:  "Test Description",
		Website_link: "https://example.com",
		Price:        100.0,
		Target_id:    1,
		Owner_id:     1,
		Clicks:       0,
		Balance:      0.0,
	}

	mock.ExpectQuery(`INSERT INTO "pad" (.+) RETURNING id;`).
		WithArgs(pad.Name, pad.Description, pad.Website_link, pad.Price, pad.Target_id, pad.Owner_id, pad.Clicks, pad.Balance).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdPad, err := repo.Create(pad)

	require.NoError(t, err)

	assert.Equal(t, 1, createdPad.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestPadRepository_Remove(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewPadRepoMock(dbInterface)
	require.NoError(t, err)

	mock.ExpectExec(`DELETE FROM "pad" WHERE id=\$1;`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove(1)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestPadRepository_Read(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewPadRepoMock(dbInterface)
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "clicks", "name", "description", "website_link", "price", "target_id", "owner_id", "balance"}).
		AddRow(1, 0, "Pad1", "Description1", "https://example1.com", 100.0, 1, 1, 0.0).
		AddRow(2, 0, "Pad2", "Description2", "https://example2.com", 200.0, 2, 2, 0.0)

	mock.ExpectQuery(`SELECT id, clicks, name, description, website_link, price, target_id, owner_id, balance FROM "pad" WHERE owner_id=\$1;`).
		WithArgs(1).
		WillReturnRows(rows)

	pads, err := repo.Read(1)

	require.NoError(t, err)

	assert.Len(t, pads, 2)

	assert.Equal(t, "Pad1", pads[0].Name)
	assert.Equal(t, "Description1", pads[0].Description)
	assert.Equal(t, "https://example1.com", pads[0].Website_link)
	assert.Equal(t, 100.0, pads[0].Price)
	assert.Equal(t, 1, pads[0].Target_id)
	assert.Equal(t, 1, pads[0].Owner_id)
	assert.Equal(t, 0.0, pads[0].Balance)
	assert.Equal(t, 1, pads[0].Id)

	assert.Equal(t, "Pad2", pads[1].Name)
	assert.Equal(t, "Description2", pads[1].Description)
	assert.Equal(t, "https://example2.com", pads[1].Website_link)
	assert.Equal(t, 200.0, pads[1].Price)
	assert.Equal(t, 2, pads[1].Target_id)
	assert.Equal(t, 2, pads[1].Owner_id)
	assert.Equal(t, 0.0, pads[1].Balance)
	assert.Equal(t, 2, pads[1].Id)

	require.NoError(t, mock.ExpectationsWereMet())
}
