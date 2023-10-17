package interfaces

import "AdHub/internal/entities"

type UserUseCase interface {
	UserRead(login string) (*entities.User, error)
	UserDelete(userMail string) error
	UserCreate(user *entities.User) (*entities.User, error)
}

type AdUseCase interface {
	AdCreate(ad *entities.Ad) (*entities.Ad, error)
	AdGetList(id int) ([]*entities.Ad, error)
}
