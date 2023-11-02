package router

import (
	"encoding/json"
	"net/http"
)

func (br *BalanceRouter) BalanceReplenishHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token  string  `json:"token"`
		Amount float64 `json:"amount"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		br.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Проверка на отрицательность (нельзя пополнить баланс на отрицательное значение)
	if request.Amount <= 0 {
		br.logger.Error("The replenishment amount must be positive")
		http.Error(w, "The amount must be positive", http.StatusBadRequest)
		return
	}

	userId, err := br.Session.GetUserId(request.Token)
	if err != nil {
		br.logger.Error("Error getting user ID from session: " + err.Error())
		http.Error(w, "Error getting user ID", http.StatusBadRequest)
		return
	}

	// Получение текущего баланса
	currentBalance, err := br.Balance.BalanceRead(userId)
	if err != nil {
		br.logger.Error("Error retrieving current balance: " + err.Error())
		http.Error(w, "Error retrieving balance", http.StatusInternalServerError)
		return
	}

	// Пополнение
	err = br.Balance.BalanceUP(request.Amount, currentBalance.Id)
	if err != nil {
		br.logger.Error("Error updating balance: " + err.Error())
		http.Error(w, "Error updating balance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // HTTP Status - 200
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentBalance)
}
