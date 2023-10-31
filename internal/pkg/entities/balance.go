package entities

type Balance struct {
	Id                int     `json:"id"`                // Id
	Total_balance     float64 `json:"total_balance"`     // Общий бюджет
	Reserved_balance  float64 `json:"reserved_balance"`  // Зарезирвированный
	Available_balance float64 `json:"available_balance"` // Свободный
}

//go:generate /Users/bincom/go/bin/mockgen -source=ad.go -destination=mock_entities/ad_mock.go
type BalanceRepoInterface interface {
	Create(s *Balance) (*Balance, error)
	Remove(id int) error
	Update(s *Balance) error
	Read(id int) (*Balance, error)
}
