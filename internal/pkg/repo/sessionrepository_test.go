package repo

import (
	"AdHub/internal/pkg/entities"
	ss "AdHub/pkg/SessionStorage"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestRedis() (*miniredis.Miniredis, *redis.Client) {
	mr, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	return mr, client
}

func TestCreateSession(t *testing.T) {
	mr, client := setupTestRedis()
	defer mr.Close()
	defer client.Close()

	mock := ss.NewMock(client)
	repo, err := NewSessionRepoMock(mock)
	require.NoError(t, err)

	session := &entities.Session{
		Token:  "token123",
		UserId: 42,
	}

	createdSession, err := repo.Create(session)

	require.NoError(t, err)
	assert.Equal(t, session.Token, createdSession.Token)
	assert.Equal(t, session.UserId, createdSession.UserId)
}

func TestReadSession(t *testing.T) {
	mr, client := setupTestRedis()
	defer mr.Close()
	defer client.Close()

	mock := ss.NewMock(client)
	repo, err := NewSessionRepoMock(mock)
	require.NoError(t, err)

	session := &entities.Session{
		Token:  "token123",
		UserId: 42,
	}

	repo.Create(session)

	retrievedSession, err := repo.Read(session)

	require.NoError(t, err)
	assert.Equal(t, session.Token, retrievedSession.Token)
	assert.Equal(t, session.UserId, retrievedSession.UserId)
}

func TestRemoveSession(t *testing.T) {
	mr, client := setupTestRedis()
	defer mr.Close()
	defer client.Close()

	mock := ss.NewMock(client)
	repo, err := NewSessionRepoMock(mock)
	require.NoError(t, err)

	session := &entities.Session{
		Token:  "token123",
		UserId: 42,
	}

	repo.Create(session)

	err = repo.Remove(session)
	require.NoError(t, err)

	exists, err := repo.Contains(session)
	require.NoError(t, err)
	assert.False(t, exists)
}

func TestContainsSession(t *testing.T) {
	mr, client := setupTestRedis()
	defer mr.Close()
	defer client.Close()

	mock := ss.NewMock(client)
	repo, err := NewSessionRepoMock(mock)
	require.NoError(t, err)

	session := &entities.Session{
		Token:  "token123",
		UserId: 42,
	}

	repo.Create(session)

	exists, err := repo.Contains(session)

	require.NoError(t, err)
	assert.True(t, exists)
}
