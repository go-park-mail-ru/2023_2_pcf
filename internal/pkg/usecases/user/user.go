package user

import (
	"AdHub/internal/pkg/entities"
)

type UserUseCase struct {
	repo entities.UserRepoInterface
}

func New(r entities.UserRepoInterface) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) UserCreate(user *entities.User) (*entities.User, error) {
	return uc.repo.Create(user)
}

func (uc *UserUseCase) UserDelete(userMail string) error {
	return uc.repo.Remove(userMail)
}

func (uc *UserUseCase) UserRead(login string) (*entities.User, error) {
	return uc.repo.Read(login)
}

func (uc *UserUseCase) UserUpdate(s *entities.User) error {
	return uc.repo.Update(s)
}
