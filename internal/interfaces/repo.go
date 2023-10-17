package interfaces

import (
	"AdHub/internal/entities"
	"database/sql"
)

type UserRepo interface {
	Create(s *entities.User) (*entities.User, error)
	Remove(mail string) error
	Get(mail string) (*sql.Rows, error)
	Update(s *entities.User) error
	Read(mail string) (*entities.User, error)
}

type AdRepo interface {
	Create(s *entities.Ad) (*entities.Ad, error)
	Remove(id int) error
	Get(id int) (*sql.Rows, error)
	GetList(id int) (*sql.Rows, error)
	Update(s *entities.Ad) error
	Read(id int) ([]*entities.Ad, error)
}
