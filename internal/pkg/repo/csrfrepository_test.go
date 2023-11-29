package repo

import (
	"AdHub/internal/pkg/entities"
	ss "AdHub/pkg/SessionStorage"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func setupTestRedis() (*miniredis.Miniredis, *redis.Client) {
	mr, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	return mr, client
}

func TestCreateCsrf(t *testing.T) {
	mr, client := setupTestRedis()
	defer mr.Close()
	defer client.Close()

	mock := ss.NewMock(client)
	repo, err := NewCsrfRepoMock(mock)
	require.NoError(t, err)

	csrf := &entities.Csrf{
		Token:  "csrfToken123",
		UserId: 42,
	}

	createdCsrf, err := repo.Create(csrf)

	require.NoError(t, err)
	assert.Equal(t, csrf.Token, createdCsrf.Token)
	assert.Equal(t, csrf.UserId, createdCsrf.UserId)
}

func TestReadCsrf(t *testing.T) {
	mr, client := setupTestRedis()
	defer mr.Close()
	defer client.Close()

	mock := ss.NewMock(client)
	repo, err := NewCsrfRepoMock(mock)
	require.NoError(t, err)

	csrf := &entities.Csrf{
		Token:  "csrfToken123",
		UserId: 42,
	}

	repo.Create(csrf)

	retrievedCsrf, err := repo.Read(csrf.UserId)

	require.NoError(t, err)
	assert.Equal(t, csrf.Token, retrievedCsrf.Token)
	assert.Equal(t, csrf.UserId, retrievedCsrf.UserId)
}

func TestRemoveCsrf(t *testing.T) {
	mr, client := setupTestRedis()
	defer mr.Close()
	defer client.Close()

	mock := ss.NewMock(client)
	repo, err := NewCsrfRepoMock(mock)
	require.NoError(t, err)

	csrf := &entities.Csrf{
		Token:  "csrfToken123",
		UserId: 42,
	}

	repo.Create(csrf)

	err = repo.Remove(csrf)
	require.NoError(t, err)
}
