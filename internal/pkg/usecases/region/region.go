package region

import (
	"AdHub/internal/pkg/entities"
)

type RegionUseCase struct {
	repo entities.RegionRepoInterface
}

func New(r entities.RegionRepoInterface) *RegionUseCase {
	return &RegionUseCase{
		repo: r,
	}
}

func (uc *RegionUseCase) RegionCreate(s *entities.Region) (*entities.Region, error) {
	return uc.repo.Create(s)
}

func (uc *RegionUseCase) RegionRemove(id int) error {
	return uc.repo.Remove(id)
}
