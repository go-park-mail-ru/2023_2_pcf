package ulink

import (
	"AdHub/internal/pkg/entities"
)

type ULinkUseCase struct {
	repo entities.ULinkRepoInterface
}

func New(r entities.ULinkRepoInterface) *ULinkUseCase {
	return &ULinkUseCase{
		repo: r,
	}
}

func (su ULinkUseCase) ULinkCreate(sr *entities.ULink) (*entities.ULink, error) {
	return su.repo.Create(sr)
}

func (su ULinkUseCase) ULinkRead(sr *entities.ULink) (*entities.ULink, error) {
	return su.repo.Read(sr)
}

func (su ULinkUseCase) ULinkRemove(sr *entities.ULink) error {
	return su.repo.Remove(sr)
}

func (su ULinkUseCase) ULinkContains(sr *entities.ULink) (bool, error) {
	return su.repo.Contains(sr)
}

func (su ULinkUseCase) GetAdId(token string) (int, error) {
	s, err := su.repo.Read(&entities.ULink{
		Token: token,
		AdId:  0,
	})
	return s.AdId, err
}
