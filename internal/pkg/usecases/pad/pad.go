package pad

import (
	"AdHub/internal/pkg/entities"
)

type PadUseCase struct {
	repo entities.PadRepoInterface
}

func New(r entities.PadRepoInterface) *PadUseCase {
	return &PadUseCase{
		repo: r,
	}
}

func (uc *PadUseCase) PadCreate(pad *entities.Pad) (*entities.Pad, error) {
	return uc.repo.Create(pad)
}

func (uc *PadUseCase) PadReadList(id int) ([]*entities.Pad, error) {
	return uc.repo.Read(id)
}

func (uc *PadUseCase) PadRead(id int) (*entities.Pad, error) {
	return uc.repo.Get(id)
}

func (uc *PadUseCase) PadRemove(id int) error {
	return uc.repo.Remove(id)
}

func (uc *PadUseCase) PadUpdate(s *entities.Pad) error {
	return uc.repo.Update(s)
}
