package user

import (
	"AdHub/internal/entities"
	"AdHub/internal/interfaces"
)

type UserUseCase struct {
	repo interfaces.UserRepo
}

func New(r interfaces.UserRepo) *UserUseCase {
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
