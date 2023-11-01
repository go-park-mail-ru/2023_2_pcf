package tag

import (
	"AdHub/internal/pkg/entities"
)

type TagUseCase struct {
	repo entities.TagRepoInterface
}

func New(r entities.TagRepoInterface) *TagUseCase {
	return &TagUseCase{
		repo: r,
	}
}

func (uc *TagUseCase) TagCreate(s *entities.Tag) (*entities.Tag, error) {
	return uc.repo.Create(s)
}

func (uc *TagUseCase) TagRemove(id int) error {
	return uc.repo.Remove(id)
}
