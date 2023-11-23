package entities

type User struct {
	Id          int    `json:"id"`         // Id
	Login       string `json:"login"`      // Логин
	Password    string `json:"password"`   // Пароль
	FName       string `json:"f_name"`     // Имя
	LName       string `json:"l_name"`     // Фамилия
	CompanyName string `json:"s_name"`     // Отчество
	Avatar      string `json:"avatar"`     // Имя файла аватара (путь или имя файла)
	BalanceId   int    `json:"balance_id"` // ID баланса
}
