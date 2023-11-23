package entities

type Pad struct {
	Id           int     `json:"id"`           // Id
	Name         string  `json:"name"`         // название объявления
	Description  string  `json:"description"`  // описание
	Website_link string  `json:"website_link"` // cайт
	Price        float64 `json:"price"`        // бюджет
	Target_id    int     `json:"target_id"`    //таргетинг
	Owner_id     int     `json:"owner_id`      // id владельца
	Clicks       int     `json:"clicks"`       // Клики
}

//go:generate /Users/bincom/go/bin/mockgen -source=pad.go -destination=mock_entities/pad_mock.go
type PadRepoInterface interface {
	Create(s *Pad) (*Pad, error)
	Remove(id int) error
	Update(s *Pad) error
	Read(id int) ([]*Pad, error)
	Get(id int) (*Pad, error)
}

type PadUseCaseInterface interface {
	PadCreate(pad *Pad) (*Pad, error)
	PadReadList(id int) ([]*Pad, error)
	PadRead(id int) (*Pad, error)
	PadRemove(id int) error
	PadUpdate(s *Pad) error
}
