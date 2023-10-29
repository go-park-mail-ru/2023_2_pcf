package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/SessionStorage"
	"context"
	"fmt"
	"strconv"
	"time"
)

type SessionRepository struct {
	store SessionStorage.SessionStorageInterface
}

func NewSessionRepo(ss SessionStorage.SessionStorageInterface) (sr SessionRepository, err error) {
	sr.store, err = ss.Open()
	return sr, err
}

func (sr SessionRepository) Create(s *entities.Session) (*entities.Session, error) {
	if len(s.Token) == 0 {
		return nil, fmt.Errorf("Token len is 0")
	}
	ctx := context.Background()
	err := sr.store.Store().Set(ctx, s.Token, s.UserId, time.Hour*24).Err()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (sr SessionRepository) Read(s *entities.Session) (*entities.Session, error) {
	if len(s.Token) == 0 {
		return nil, fmt.Errorf("Token len is 0")
	}

	ctx := context.Background()
	userId, err := sr.store.Store().Get(ctx, s.Token).Result()
	if err != nil {
		return nil, err
	}

	s.UserId, err = strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (sr SessionRepository) Remove(s *entities.Session) error {
	if len(s.Token) == 0 {
		return fmt.Errorf("Token len is 0")
	}

	ctx := context.Background()
	sr.store.Store().Del(ctx, s.Token)
	return nil
}

func (sr SessionRepository) Contains(s *entities.Session) (bool, error) {
	ctx := context.Background()
	exists, err := sr.store.Store().Exists(ctx, s.Token).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}
