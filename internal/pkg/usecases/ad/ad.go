package ad

import (
	"AdHub/internal/pkg/entities"
)

type AdUseCase struct {
	repo entities.AdRepoInterface
}

func New(r entities.AdRepoInterface) *AdUseCase {
	return &AdUseCase{
		repo: r,
	}
}

func (uc *AdUseCase) AdCreate(ad *entities.Ad) (*entities.Ad, error) {
	return uc.repo.Create(ad)
}

func (uc *AdUseCase) AdReadList(id int) ([]*entities.Ad, error) {
	return uc.repo.Read(id)
}
