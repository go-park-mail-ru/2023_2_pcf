package ad

import (
	"AdHub/internal/entities"
	"AdHub/internal/interfaces"
)

type AdUseCase struct {
	repo interfaces.AdRepo
}

func New(r interfaces.AdRepo) *AdUseCase {
	return &AdUseCase{
		repo: r,
	}
}

func (uc *AdUseCase) AdCreate(ad *entities.Ad) (*entities.Ad, error) {
	return uc.repo.Create(ad)
}

func (uc *AdUseCase) AdGetList(id int) ([]*entities.Ad, error) {
	return uc.repo.Read(id)
}
