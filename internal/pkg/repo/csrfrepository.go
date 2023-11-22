package repo

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/CsrfStorage"
	"context"
	"fmt"
	"strconv"
	"time"
)

type CsrfRepository struct {
	store CsrfStorage.CsrfStorageInterface
}

func NewCsrfRepo(ss CsrfStorage.CsrfStorageInterface) (cr CsrfRepository, err error) {
	cr.store, err = ss.Open()
	return cr, err
}

func (cr CsrfRepository) Create(c *entities.Csrf) (*entities.Csrf, error) {
	if len(c.Token) == 0 {
		return nil, fmt.Errorf("CSRF token length is 0")
	}
	ctx := context.Background()
	err := cr.store.Store().Set(ctx, strconv.Itoa(c.UserId), c.Token, time.Hour*24).Err()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cr CsrfRepository) Read(userId int) (*entities.Csrf, error) {
	ctx := context.Background()
	//ключ - это ID пользователя
	csrfToken, err := cr.store.Store().Get(ctx, strconv.Itoa(userId)).Result()
	if err != nil {
		return nil, err
	}

	return &entities.Csrf{
		UserId: userId,
		Token:  csrfToken,
	}, nil
}

func (cr CsrfRepository) Remove(c *entities.Csrf) error {
	if len(c.Token) == 0 {
		return fmt.Errorf("CSRF token length is 0")
	}

	ctx := context.Background()
	cr.store.Store().Del(ctx, c.Token)
	return nil
}