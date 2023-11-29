package repo

import (
	"AdHub/internal/pkg/entities"
	pg "AdHub/pkg/db"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPadCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &PadRepository{store: pg.NewMock(db)}

	pad := &entities.Pad{
		Name:         "Test Pad",
		Description:  "Test Description",
		Website_link: "https://example.com",
		Price:        100.0,
		Target_id:    1,
		Owner_id:     1,
		Clicks:       0,
		Balance:      0,
	}

	mock.ExpectQuery(`INSERT INTO "pad" (.+) RETURNING id;`).
		WithArgs(pad.Name, pad.Description, pad.Website_link, pad.Price, pad.Target_id, pad.Owner_id, pad.Clicks, pad.Balance).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdPad, err := repo.Create(pad)

	require.NoError(t, err)
	assert.Equal(t, 1, createdPad.Id)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestPadRemove(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &PadRepository{store: pg.NewMock(db)}

	mock.ExpectExec(`DELETE FROM "pad" WHERE id=\$1;`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove(1)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestPadUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &PadRepository{store: pg.NewMock(db)}

	pad := &entities.Pad{
		Id:           1,
		Name:         "Updated Pad",
		Description:  "Updated Description",
		Website_link: "https://updated.com",
		Price:        200.0,
		Target_id:    2,
		Owner_id:     1,
		Clicks:       10,
		Balance:      1,
	}

	mock.ExpectExec(`UPDATE "pad" SET (.+) WHERE id=\$9;`).
		WithArgs(pad.Name, pad.Description, pad.Website_link, pad.Price, pad.Target_id, pad.Owner_id, pad.Clicks, float64(pad.Balance), pad.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(pad)

	require.NoError(t, nil)
}

func TestPadRead(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &PadRepository{store: pg.NewMock(db)}

	rows := sqlmock.NewRows([]string{"id", "clicks", "name", "description", "website_link", "price", "target_id", "owner_id", "balance"}).
		AddRow(1, 10, "Pad1", "Description1", "https://example1.com", 100.0, 1, 1, 50).
		AddRow(2, 20, "Pad2", "Description2", "https://example2.com", 200.0, 2, 2, 100)

	mock.ExpectQuery(`SELECT (.+) FROM "pad" WHERE owner_id=\$1;`).
		WithArgs(1).
		WillReturnRows(rows)

	pads, err := repo.Read(1)

	require.NoError(t, err)
	assert.Len(t, pads, 2)
	assert.Equal(t, "Pad1", pads[0].Name)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestPadGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &PadRepository{store: pg.NewMock(db)}

	rows := sqlmock.NewRows([]string{"id", "clicks", "name", "description", "website_link", "price", "target_id", "owner_id", "balance"}).
		AddRow(1, 10, "Pad1", "Description1", "https://example1.com", 100.0, 1, 1, 50)

	mock.ExpectQuery(`SELECT (.+) FROM "pad" WHERE id=\$1;`).
		WithArgs(1).
		WillReturnRows(rows)

	pad, err := repo.Get(1)

	require.NoError(t, err)
	assert.NotNil(t, pad)
	assert.Equal(t, "Pad1", pad.Name)

	require.NoError(t, mock.ExpectationsWereMet())
}
