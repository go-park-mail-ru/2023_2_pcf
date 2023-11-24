package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/SessionStorage"
	"context"
	"fmt"
	"strconv"
	"time"
)

type ULinkRepository struct {
	store SessionStorage.SessionStorageInterface
}

func NewULinkRepoMock(ss SessionStorage.SessionStorageInterface) (*ULinkRepository, error) {
	sr := &ULinkRepository{
		store: ss,
	}
	return sr, nil
}

func NewULinkRepo(ss SessionStorage.SessionStorageInterface) (sr ULinkRepository, err error) {
	sr.store, err = ss.Open()
	return sr, err
}

func (sr ULinkRepository) Create(s *entities.ULink) (*entities.ULink, error) {
	if len(s.Token) == 0 {
		return nil, fmt.Errorf("Token len is 0")
	}
	ctx := context.Background()
	err := sr.store.Store().Set(ctx, s.Token, s.AdId, time.Hour*24).Err()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (sr ULinkRepository) Read(s *entities.ULink) (*entities.ULink, error) {
	if len(s.Token) == 0 {
		return nil, fmt.Errorf("Token len is 0")
	}

	ctx := context.Background()
	AdId, err := sr.store.Store().Get(ctx, s.Token).Result()
	if err != nil {
		return nil, err
	}

	s.AdId, err = strconv.Atoi(AdId)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (sr ULinkRepository) Remove(s *entities.ULink) error {
	if len(s.Token) == 0 {
		return fmt.Errorf("Token len is 0")
	}

	ctx := context.Background()
	sr.store.Store().Del(ctx, s.Token)
	return nil
}

func (sr ULinkRepository) Contains(s *entities.ULink) (bool, error) {
	ctx := context.Background()
	exists, err := sr.store.Store().Exists(ctx, s.Token).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}
