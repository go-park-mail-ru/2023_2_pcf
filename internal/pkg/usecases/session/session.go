package session

import (
	"AdHub/internal/pkg/entities"
)

type SessionUseCase struct {
	repo entities.SessionRepoInterface
}

func New(r entities.SessionRepoInterface) *SessionUseCase {
	return &SessionUseCase{
		repo: r,
	}
}

func (su SessionUseCase) SessionCreate(sr *entities.Session) (*entities.Session, error) {
	return su.repo.Create(sr)
}

func (su SessionUseCase) SessionRead(sr *entities.Session) (*entities.Session, error) {
	return su.repo.Read(sr)
}

func (su SessionUseCase) SessionRemove(sr *entities.Session) error {
	return su.repo.Remove(sr)
}

func (su SessionUseCase) SessionContains(sr *entities.Session) (bool, error) {
	return su.repo.Contains(sr)
}
