package entities

type Target struct {
	Id        int      `json:"id"`        // Id
	Name      string   `json:"name"`      // Название
	Owner_id  int      `json:"owner_id"`  // Владелец
	Gender    string   `json:"gender"`    // Пол
	Min_age   int      `json:"min_age"`   // Минимальный возраст
	Max_age   int      `json:"max_age"`   // Максимальный возраст
	Interests []string `json:"interests"` // Интересы
	Tags      []string `json:"tags"`      // Тэги
	Keys      []string `json:"keys"`      // Ключи
	Regions   []string `json:"regions"`   // Регионы
}

//go:generate /Users/bincom/go/bin/mockgen -source=target.go -destination=mock_entities/target_mock.go
type TargetRepoInterface interface {
	Create(s *Target) (*Target, error)
	Remove(id int) error
	Update(s *Target) error
	Read(id int) (*Target, error)
	ReadList(id int) ([]*Target, error)
	ReadRandom() ([]*Target, error)
}

type TargetUseCaseInterface interface {
	TargetCreate(s *Target) (*Target, error)
	TargetRead(id int) (*Target, error)
	TargetReadList(id int) ([]*Target, error)
	TargetRemove(id int) error
	TargetUpdate(s *Target) error
	TargetRandom() ([]*Target, error)
}
