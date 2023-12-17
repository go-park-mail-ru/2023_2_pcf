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

func (uc *AdUseCase) AdRead(id int) (*entities.Ad, error) {
	return uc.repo.Get(id)
}

func (uc *AdUseCase) AdRemove(id int) error {
	return uc.repo.Remove(id)
}

func (uc *AdUseCase) AdUpdate(s *entities.Ad) error {
	return uc.repo.Update(s)
}

func (uc *AdUseCase) AdByTarget(id int) ([]*entities.Ad, error) {
	return uc.repo.ReaByTarget(id)
}
