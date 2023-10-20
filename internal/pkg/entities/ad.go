package entities

type Ad struct {
	Id          int    `json:"id"`          // Id
	Name        string `json:"name"`        // название объявления
	Description string `json:"description"` // описание
	Sector      string `json:"sector"`      // cфера
	Owner_id    int    `json:"owner_id`     // id владельца
}

type AdRepoInterface interface {
	Create(s *Ad) (*Ad, error)
	Remove(id int) error
	Update(s *Ad) error
	Read(id int) ([]*Ad, error)
}

type AdUseCaseInterface interface {
	AdCreate(ad *Ad) (*Ad, error)
	AdReadList(id int) ([]*Ad, error)
}
