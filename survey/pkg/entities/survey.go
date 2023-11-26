package entities

type Survey struct {
	Id       int    `json:"id"`       // Id
	Question string `json:"question"` // Вопрос
	Type     int    `json:"type"`     // Тип
}

//go:generate /Users/bincom/go/bin/mockgen -source=target.go -destination=mock_entities/target_mock.go
type SurveyRepoInterface interface {
	Create(s *Survey) (*Survey, error)
	Remove(id int) error
	Update(s *Survey) error
	Read(id int) (*Survey, error)
	ReadList() (*[]Survey, error)
}

type SurveyUseCaseInterface interface {
	SurveyCreate(s *Survey) (*Survey, error)
	SurveyRead(id int) (*Survey, error)
	SurveyRemove(id int) error
	SurveyUpdate(s *Survey) error
	SurveyList() (*[]Survey, error)
}
