package entities

type Region struct {
	Id       int    `json:"id"`        // Id
	Name     string `json:"name"`      // Название
	TargetID int    `json:"target_id"` //Id таргета
}

//go:generate /Users/bincom/go/bin/mockgen -source=region.go -destination=mock_entities/region_mock.go
type RegionRepoInterface interface {
	Create(s *Region) (*Region, error)
	Remove(id int) error
	Update(s *Region) error
	Read(id int) (*Region, error)
}

type RegionUseCaseInterface interface {
	RegionCreate(s *Interest) (*Interest, error)
	RegionRemove(id int) error
}
