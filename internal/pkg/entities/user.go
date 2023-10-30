package entities

import (
	"strings"
)

const (
	MIN_PASSWORD_LEN = 5
	MIN_EMAIL_LEN    = 5 // q@w.e
)

// Примерная структура модели для БД
type User struct {
	Id       int    `json:"id"`         // Id
	Login    string `json:"login"`      // Логин
	Password string `json:"password"`   // Пароль
	FName    string `json:"first_name"` // Имя
	LName    string `json:"last_name"`  // Фамилия
}

//go:generate /Users/bincom/go/bin/mockgen -source=user.go -destination=mock_entities/user_mock.go
type UserRepoInterface interface {
	Create(s *User) (*User, error)
	Remove(mail string) error
	Update(s *User) error
	Read(mail string) (*User, error)
}

type UserUseCaseInterface interface {
	UserRead(login string) (*User, error)
	UserDelete(userMail string) error
	UserCreate(user *User) (*User, error)
}

func (user *User) ValidateEmail() bool {
	return len(user.Login) >= MIN_EMAIL_LEN && strings.Contains(user.Login, "@") && strings.Contains(user.Login, ".")
}

func (user *User) ValidatePassword() bool {
	return len(user.Password) >= MIN_PASSWORD_LEN
}

func (user *User) ValidateFName() bool {
	return len(user.FName) > 0
}

func (user *User) ValidateLName() bool {
	return len(user.LName) > 0
}
