package tag

import (
	"AdHub/internal/pkg/entities"
)

type TargetUseCase struct {
	repo entities.TargetRepoInterface
}

func New(r entities.TagRepoInterface) *TargetUseCase {
	return &TargetUseCase{
		repo: r,
	}
}

type TargetRepoInterface interface {
	Create(s *Target) (*Target, error)
	Remove(id int) error
	Update(s *Target) error
	Read(id int) (*Target, error)
	GetTargetInterests(targetID int) ([]Interest, error)
	GetTargetRegions(targetID int) ([]Region, error)
	GetTargetTags(targetID int) ([]Tag, error)
}

type TargetUseCaseInterface interface {
	TargetCreate(s *Target) (*Target, error)
	TargetRead(id int) (*Target, error)
	TargetRemove(id int) error
	TargetUpdate(s *Target) error
}
