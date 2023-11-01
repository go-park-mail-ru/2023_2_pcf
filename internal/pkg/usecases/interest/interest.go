package interest

import (
	"AdHub/internal/pkg/entities"
)

type InterestUseCase struct {
	repo entities.InterestRepoInterface
}

func New(r entities.InterestRepoInterface) *InterestUseCase {
	return &InterestUseCase{
		repo: r,
	}
}

func (uc *InterestUseCase) InterestCreate(s *entities.Interest) (*entities.Interest, error) {
	return uc.repo.Create(s)
}

func (uc *InterestUseCase) InterestRemove(id int) error {
	return uc.repo.Remove(id)
}
