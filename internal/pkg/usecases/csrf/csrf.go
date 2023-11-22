package csrf

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/cryptoUtils"
)

type CsrfUseCase struct {
	repo entities.CsrfRepoInterface
}

func NewCsrfUseCase(r entities.CsrfRepoInterface) *CsrfUseCase {
	return &CsrfUseCase{
		repo: r,
	}
}

func (cu *CsrfUseCase) CsrfCreate(userId int) (*entities.Csrf, error) {
	var tokenLen = 32
	newCsrf := &entities.Csrf{UserId: userId}

	var err error
	newCsrf.Token, err = cryptoUtils.GenToken(tokenLen)
	if err != nil {
		return nil, err
	}

	newCsrf, err = cu.repo.Create(newCsrf)
	if err != nil {
		return nil, err
	}

	return newCsrf, nil
}

func (cu *CsrfUseCase) CsrfRemove(sr *entities.Csrf) error {
	return cu.repo.Remove(sr)
}

func (cu *CsrfUseCase) GetByUserId(userId int) (*entities.Csrf, error) {
	return cu.repo.Read(userId)
}
