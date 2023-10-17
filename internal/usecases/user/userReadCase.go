package user

import (
	"AdHub/internal/entities"
	"AdHub/internal/interfaces"
)

func UserRead(repo interfaces.UserRepo, login string) (*entities.User, error) {
	return .Read(login)
}
