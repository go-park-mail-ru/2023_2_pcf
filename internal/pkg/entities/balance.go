package entities

type Balance struct {
	Id                int     `json:"id"`                // Id
	Total_balance     float64 `json:"total_balance"`     // Общий бюджет
	Reserved_balance  float64 `json:"reserved_balance"`  // Зарезирвированный
	Available_balance float64 `json:"available_balance"` // Свободный
}

//go:generate /Users/bincom/go/bin/mockgen -source=balance.go -destination=mock_entities/balance_mock.go
type BalanceRepoInterface interface {
	Create(s *Balance) (*Balance, error)
	Remove(id int) error
	Update(s *Balance) error
	Read(id int) (*Balance, error)
}

type BalanceUseCaseInterface interface {
	BalanceCreate(s *Balance) (*Balance, error)
	BalanceRead(id int) (*Balance, error)
	BalanceRemove(id int) error
	BalanceUP(sum float64, id int) error
	BalanceDown(sum float64, id int) error
	BalanceReserve(sum float64, id int) error
}
