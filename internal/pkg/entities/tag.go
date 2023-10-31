package entities

type Tag struct {
	Id   int    `json:"id"`   // Id
	Name string `json:"name"` // Название
}

//go:generate /Users/bincom/go/bin/mockgen -source=tag.go -destination=mock_entities/tag_mock.go
type TagRepoInterface interface {
	Create(s *Tag) (*Tag, error)
	Remove(id int) error
	Update(s *Tag) error
	Read(id int) (*Tag, error)
}
