package entities

type Rate struct {
	Id        int `json:"id"` // Id
	User_id   int `json:"user_id"`
	Rate      int `json:"rate"`
	Survey_id int `json:"survey_id"`
}

type Response struct {
	Count     int     `json:"count"`
	Avg_rate  float64 `json:"avg_rate"`
	Survey_id int     `json:"survey_id"`
}

//go:generate /Users/bincom/go/bin/mockgen -source=target.go -destination=mock_entities/target_mock.go
type RateRepoInterface interface {
	Create(s *Rate) (*Rate, error)
	Read(id int) ([]*Rate, error)
}

type RateUseCaseInterface interface {
	RateCreate(s *Rate) (*Rate, error)
	RateRead(id int) (*Response, error)
}
