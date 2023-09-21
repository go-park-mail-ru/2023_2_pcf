package model

// Примерная структура модели для БД
type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
