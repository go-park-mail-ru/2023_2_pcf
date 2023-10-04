package models

type Ad struct {
	Id          int    `json:"id"`          // Id
	Name        string `json:"name"`        // название объявления
	Description string `json:"description"` // описание
	Sector      string `json:"sector"`      // cфера
	Owner_id    int    `json:"owner_id`     // id владельца
}
