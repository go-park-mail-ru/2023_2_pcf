package entities

type Interest struct {
	Id   int    `json:"id"`   // Id
	Name string `json:"name"` // Название
}

//go:generate /Users/bincom/go/bin/mockgen -source=interest.go -destination=mock_entities/interest_mock.go
type InterestRepoInterface interface {
	Create(s *Interest) (*Interest, error)
	Remove(id int) error
	Update(s *Interest) error
	Read(id int) (*Interest, error)
}
