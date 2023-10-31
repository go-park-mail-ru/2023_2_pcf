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

func TestCreateTag(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTagRepoMock(dbInterface)
	require.NoError(t, err)

	tag := &entities.Tag{
		Name: "Test Tag",
	}

	mock.ExpectQuery(`INSERT INTO "tags" (.+) RETURNING id;`).
		WithArgs(tag.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdTag, err := repo.Create(tag)

	require.NoError(t, err)

	assert.Equal(t, 1, createdTag.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRemoveTag(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTagRepoMock(dbInterface)
	require.NoError(t, err)

	mock.ExpectExec(`DELETE FROM "tags" WHERE id=\$1;`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove(1)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateTag(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTagRepoMock(dbInterface)
	require.NoError(t, err)

	tag := &entities.Tag{
		Id:   1,
		Name: "Updated Tag",
	}

	mock.ExpectExec(`UPDATE "tags" SET name=\$1 WHERE id=\$2;`).
		WithArgs(tag.Name, tag.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(tag)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReadTag(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTagRepoMock(dbInterface)
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Test Tag")

	mock.ExpectQuery(`SELECT id, name FROM "tags" WHERE id=\$1;`).
		WillReturnRows(rows)

	tag, err := repo.Read(1)

	require.NoError(t, err)

	assert.Equal(t, 1, tag.Id)
	assert.Equal(t, "Test Tag", tag.Name)

	require.NoError(t, mock.ExpectationsWereMet())
}
