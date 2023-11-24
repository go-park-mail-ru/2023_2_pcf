package target

import (
	"AdHub/internal/pkg/entities"
)

type TargetUseCase struct {
	repo entities.TargetRepoInterface
}

func New(r entities.TargetRepoInterface) *TargetUseCase {
	return &TargetUseCase{
		repo: r,
	}
}

func (uc *TargetUseCase) TargetCreate(s *entities.Target) (*entities.Target, error) {
	return uc.repo.Create(s)
}

func (uc *TargetUseCase) TargetRemove(id int) error {
	return uc.repo.Remove(id)
}

func (uc *TargetUseCase) TargetRead(id int) (*entities.Target, error) {
	return uc.repo.Read(id)
}

func (uc *TargetUseCase) TargetReadList(id int) ([]*entities.Target, error) {
	return uc.repo.ReadList(id)
}

func (uc *TargetUseCase) TargetUpdate(s *entities.Target) error {
	return uc.repo.Update(s)
}

func (uc *TargetUseCase) TargetRandom() ([]*entities.Target, error) {
	return uc.repo.ReadRandom()
}
