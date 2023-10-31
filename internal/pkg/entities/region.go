package entities

type Region struct {
	Id   int    `json:"id"`   // Id
	Name string `json:"name"` // Название
}

//go:generate /Users/bincom/go/bin/mockgen -source=region.go -destination=mock_entities/region_mock.go
type RegionRepoInterface interface {
	Create(s *Region) (*Region, error)
	Remove(id int) error
	Update(s *Region) error
	Read(id int) (*Region, error)
}
