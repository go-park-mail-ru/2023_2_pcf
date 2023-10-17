package interfaces

import (
	"AdHub/internal/entities"
	"database/sql"
)

type UserRepo interface {
	Configure(DB Db) (*UserRepo, error)
	Create(s *entities.User) (*entities.User, error)
	Remove(mail string) error
	Get(mail string) (*sql.Rows, error)
	Update(s *entities.User) error
	Read(mail string) (*entities.User, error)
}
