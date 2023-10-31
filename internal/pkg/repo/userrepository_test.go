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

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewUserRepoMock(dbInterface)
	require.NoError(t, err)

	user := &entities.User{
		Login:    "testuser",
		Password: "password",
		FName:    "John",
		LName:    "Doe",
	}

	mock.ExpectQuery("INSERT INTO \"user\" (.+) RETURNING id;").
		WithArgs(user.Login, user.Password, user.FName, user.LName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdUser, err := repo.Create(user)

	require.NoError(t, err)

	assert.Equal(t, 1, createdUser.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRemoveUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewUserRepoMock(dbInterface)
	require.NoError(t, err)

	mock.ExpectExec(`DELETE FROM "user" WHERE login=\$1;`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove("testuser")

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewUserRepoMock(dbInterface)
	require.NoError(t, err)

	user := &entities.User{
		Id:       1,
		Login:    "updateduser",
		Password: "newpassword",
		FName:    "Jane",
		LName:    "Smith",
	}

	mock.ExpectExec(`UPDATE "user" SET login=\$1, password=\$2, f_name=\$3, l_name=\$4 WHERE id=\$5;`).
		WithArgs(user.Login, user.Password, user.FName, user.LName, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(user)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReadUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewUserRepoMock(dbInterface)
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "login", "password", "f_name", "l_name"}).
		AddRow(1, "testuser", "password", "John", "Doe")

	mock.ExpectQuery(`SELECT id, login, password, f_name, l_name FROM "user" WHERE login=\$1;`).
		WillReturnRows(rows)

	user, err := repo.Read("testuser")

	require.NoError(t, err)

	assert.Equal(t, "testuser", user.Login)
	assert.Equal(t, "password", user.Password)
	assert.Equal(t, "John", user.FName)
	assert.Equal(t, "Doe", user.LName)

	require.NoError(t, mock.ExpectationsWereMet())
}