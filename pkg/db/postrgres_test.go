package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPg_Open(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	pg := &Pg{db: db}

	_, err = pg.Open()
	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %s", err)
	}
}

func TestPg_Close(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	pg := &Pg{db: db}

	mock.ExpectClose()

	pg.Close()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %s", err)
	}
}
