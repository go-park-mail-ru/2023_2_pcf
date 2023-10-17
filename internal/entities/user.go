package entities

import "strings"

const (
	MIN_PASSWORD_LEN = 5
	MIN_EMAIL_LEN    = 5 // q@w.e
)

// Примерная структура модели для БД
type User struct {
	Id       int    `json:"id"`       // Id
	Login    string `json:"login"`    // Логин
	Password string `json:"password"` // Пароль
	FName    string `json:"f_name"`   // Имя
	LName    string `json:"l_name"`   // Фамилия
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
