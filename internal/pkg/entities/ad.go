package entities

type Ad struct {
	Id           int     `json:"id"`           // Id
	Name         string  `json:"name"`         // название объявления
	Description  string  `json:"description"`  // описание
	Website_link string  `json:"website_link"` // cайт
	Budget       float64 `json:"budget"`       // бюджет
	Target_id    int     `json:"target_id"`    //таргетинг
	Image_link   string  `json:"image_link"`   //баннер(картинка)
	Owner_id     int     `json:"owner_id`      // id владельца

}

//go:generate /Users/bincom/go/bin/mockgen -source=ad.go -destination=mock_entities/ad_mock.go
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
