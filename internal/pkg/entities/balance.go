package entities

type Balance struct {
	Id                int     `json:"id"`                // Id
	Total_balance     float64 `json:"total_balance"`     // Общий бюджет
	Reserved_balance  float64 `json:"reserved_balance"`  // Зарезирвированный
	Available_balance float64 `json:"available_balance"` // Свободный
}
