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
		Name:      "Test Target",
		Owner_id:  1,
		Gender:    "Male",
		Min_age:   18,
		Max_age:   50,
		Tags:      []string{"tag1", "tag2"},
		Regions:   []string{"region1", "region2"},
		Interests: []string{"interest1", "interest2"},
		Keys:      []string{"key1", "key2"},
	}

	mock.ExpectQuery(`INSERT INTO "target" (.+) RETURNING id;`).
		WithArgs(target.Name, target.Owner_id, target.Gender, target.Min_age, target.Max_age, "tag1, tag2", "region1, region2", "interest1, interest2", "key1, key2").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdTarget, err := repo.Create(target)

	require.NoError(t, err)

	assert.Equal(t, 1, createdTarget.Id)

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
		Id:        1,
		Name:      "Updated Target",
		Owner_id:  2,
		Gender:    "Female",
		Min_age:   20,
		Max_age:   60,
		Tags:      []string{"tag3", "tag4"},
		Regions:   []string{"region3", "region4"},
		Interests: []string{"interest3", "interest4"},
		Keys:      []string{"key3", "key4"},
	}

	mock.ExpectExec(`UPDATE "target" SET name=\$1, owner_id=\$2, gender=\$3, min_age=\$4, max_age=\$5, tags=\$6, regions=\$7, interests=\$8, keys=\$9 WHERE id=\$10;`).
		WithArgs(target.Name, target.Owner_id, target.Gender, target.Min_age, target.Max_age, "tag3, tag4", "region3, region4", "interest3, interest4", "key3, key4", target.Id).
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

	rows := sqlmock.NewRows([]string{"id", "name", "owner_id", "gender", "min_age", "max_age", "tags", "regions", "interests", "keys"}).
		AddRow(1, "Test Target", 1, "Male", 18, 50, "tag1, tag2", "region1, region2", "interest1, interest2", "key1, key2")

	mock.ExpectQuery(`SELECT id, name, owner_id, gender, min_age, max_age, tags, regions, interests, keys FROM "target" WHERE id=\$1;`).
		WillReturnRows(rows)

	target, err := repo.Read(1)

	require.NoError(t, err)

	assert.Equal(t, 1, target.Id)
	assert.Equal(t, "Test Target", target.Name)
	assert.Equal(t, 1, target.Owner_id)
	assert.Equal(t, "Male", target.Gender)
	assert.Equal(t, 18, target.Min_age)
	assert.Equal(t, 50, target.Max_age)
	assert.Len(t, target.Tags, 2)
	assert.Len(t, target.Regions, 2)
	assert.Len(t, target.Interests, 2)
	assert.Len(t, target.Keys, 2)

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

func TestReadListTarget(t *testing.T) {
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

	rows := sqlmock.NewRows([]string{"id", "name", "owner_id", "gender", "min_age", "max_age", "tags", "regions", "interests", "keys"}).
		AddRow(1, "Test Target 1", 1, "Male", 18, 50, "tag1, tag2", "region1, region2", "interest1, interest2", "key1, key2").
		AddRow(2, "Test Target 2", 1, "Female", 20, 60, "tag3, tag4", "region3, region4", "interest3, interest4", "key3, key4")

	mock.ExpectQuery(`SELECT id, name, owner_id, gender, min_age, max_age, tags, regions, interests, keys FROM "target" WHERE owner_id=\$1;`).
		WithArgs(1).
		WillReturnRows(rows)

	targets, err := repo.ReadList(1)

	require.NoError(t, err)
	require.Len(t, targets, 2)

	target1 := targets[0]
	assert.Equal(t, 1, target1.Id)
	assert.Equal(t, "Test Target 1", target1.Name)

	target2 := targets[1]
	assert.Equal(t, 2, target2.Id)
	assert.Equal(t, "Test Target 2", target2.Name)

	require.NoError(t, mock.ExpectationsWereMet())
}
