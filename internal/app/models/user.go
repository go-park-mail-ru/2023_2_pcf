package models

// Примерная структура модели для БД
type User struct {
	Id       int    `json:"id"`       // Id
	Login    string `json:"login"`    // Логин
	Password string `json:"password"` // Пароль
	FName    string `json:"f_name"`   // Имя
	LName    string `json:"l_name"`   // Фамилия
}
