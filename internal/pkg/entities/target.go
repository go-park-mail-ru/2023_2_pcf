package entities

type Target struct {
	Id        int        `json:"id"`        // Id
	Name      string     `json:"name"`      // Название
	Owner_id  int        `json:"owner_id"`  // Владелец
	Gender    string     `json:"gender"`    // Пол
	Min_age   int        `json:"min_age"`   // Минимальный возраст
	Max_age   int        `json:"max_age"`   // Максимальный возраст
	Interests []Interest `json:"interests"` // Интересы
	Tags      []Tag      `json:"tags"`      // Тэги
	Regions   []Region   `json:"regions"`   // Регионы
}
