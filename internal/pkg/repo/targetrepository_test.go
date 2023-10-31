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

func TestCreateTarget(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTargetRepoMock(dbInterface)
	require.NoError(t, err)

	target := &entities.Target{
		Name:     "Test Target",
		Owner_id: 1,
		Gender:   "Male",
		Min_age:  18,
		Max_age:  50,
	}

	mock.ExpectQuery(`INSERT INTO "targets" (.+) RETURNING id;`).
		WithArgs(target.Name, target.Owner_id, target.Gender, target.Min_age, target.Max_age).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdTarget, err := repo.Create(target)

	require.NoError(t, err)

	assert.Equal(t, 1, createdTarget.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRemoveTarget(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTargetRepoMock(dbInterface)
	require.NoError(t, err)

	mock.ExpectExec(`DELETE FROM "targets" WHERE id=\$1;`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove(1)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateTarget(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTargetRepoMock(dbInterface)
	require.NoError(t, err)

	target := &entities.Target{
		Id:       1,
		Name:     "Updated Target",
		Owner_id: 2,
		Gender:   "Female",
		Min_age:  20,
		Max_age:  60,
	}

	mock.ExpectExec(`UPDATE "targets" SET name=\$1, owner_id=\$2, gender=\$3, min_age=\$4, max_age=\$5 WHERE id=\$6;`).
		WithArgs(target.Name, target.Owner_id, target.Gender, target.Min_age, target.Max_age, target.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(target)

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReadTarget(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTargetRepoMock(dbInterface)
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "name", "owner_id", "gender", "min_age", "max_age"}).
		AddRow(1, "Test Target", 1, "Male", 18, 50)

	mock.ExpectQuery(`SELECT id, name, owner_id, gender, min_age, max_age FROM "targets" WHERE id=\$1;`).
		WillReturnRows(rows)

	target, err := repo.Read(1)

	require.NoError(t, err)

	assert.Equal(t, 1, target.Id)
	assert.Equal(t, "Test Target", target.Name)
	assert.Equal(t, 1, target.Owner_id)
	assert.Equal(t, "Male", target.Gender)
	assert.Equal(t, 18, target.Min_age)
	assert.Equal(t, 50, target.Max_age)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTargetTags(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTargetRepoMock(dbInterface)
	require.NoError(t, err)

	targetID := 1

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Tag1").
		AddRow(2, "Tag2")

	mock.ExpectQuery(`
        SELECT t.id, t.name
        FROM "tags" t
        JOIN "target_tags" tt ON t.id = tt.tag_id
        WHERE tt.target_id = \$1;
    `).
		WithArgs(targetID).
		WillReturnRows(rows)

	tags, err := repo.GetTargetTags(targetID)

	require.NoError(t, err)
	assert.Len(t, tags, 2)
	assert.Equal(t, 1, tags[0].Id)
	assert.Equal(t, "Tag1", tags[0].Name)
	assert.Equal(t, 2, tags[1].Id)
	assert.Equal(t, "Tag2", tags[1].Name)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTargetRegions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTargetRepoMock(dbInterface)
	require.NoError(t, err)

	targetID := 1

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Region1").
		AddRow(2, "Region2")

	mock.ExpectQuery(`
        SELECT r.id, r.name
        FROM "regions" r
        JOIN "target_regions" tr ON r.id = tr.region_id
        WHERE tr.target_id = \$1;
    `).
		WithArgs(targetID).
		WillReturnRows(rows)

	regions, err := repo.GetTargetRegions(targetID)

	require.NoError(t, err)
	assert.Len(t, regions, 2)
	assert.Equal(t, 1, regions[0].Id)
	assert.Equal(t, "Region1", regions[0].Name)
	assert.Equal(t, 2, regions[1].Id)
	assert.Equal(t, "Region2", regions[1].Name)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTargetInterests(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbInterface := pg.NewMock(db)
	repo, err := NewTargetRepoMock(dbInterface)
	require.NoError(t, err)

	targetID := 1

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Interest1").
		AddRow(2, "Interest2")

	mock.ExpectQuery(`
        SELECT i.id, i.name
        FROM "interests" i
        JOIN "target_interests" ti ON i.id = ti.interest_id
        WHERE ti.target_id = \$1;
    `).
		WithArgs(targetID).
		WillReturnRows(rows)

	interests, err := repo.GetTargetInterests(targetID)

	require.NoError(t, err)
	assert.Len(t, interests, 2)
	assert.Equal(t, 1, interests[0].Id)
	assert.Equal(t, "Interest1", interests[0].Name)
	assert.Equal(t, 2, interests[1].Id)
	assert.Equal(t, "Interest2", interests[1].Name)

	require.NoError(t, mock.ExpectationsWereMet())
}
