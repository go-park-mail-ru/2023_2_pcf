package router

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (br *BalanceRouter) BalanceReplenishHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Amount string `json:"amount"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		br.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	amount, err := strconv.ParseFloat(request.Amount, 64)
	if err != nil {
		br.logger.Error("Invalid amount: " + err.Error())
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// Проверка на отрицательность (нельзя пополнить баланс на отрицательное значение)
	if amount <= 0 {
		br.logger.Error("The replenishment amount must be positive")
		http.Error(w, "The amount must be positive", http.StatusBadRequest)
		return
	}

	uidAny := r.Context().Value("userId")
	userId, ok := uidAny.(int)
	if !ok {
		br.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	// Получение текущего баланса
	user, err := br.User.UserReadById(userId)
	if err != nil {
		br.logger.Error("Error retrieving user: " + err.Error())
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}
	currentBalance, err := br.Balance.BalanceRead(user.BalanceId)
	if err != nil {
		br.logger.Error("Error retrieving current balance: " + err.Error())
		http.Error(w, "Error retrieving balance", http.StatusInternalServerError)
		return
	}

	// Пополнение
	err = br.Balance.BalanceUP(amount, currentBalance.Id)
	if err != nil {
		br.logger.Error("Error updating balance: " + err.Error())
		http.Error(w, "Error updating balance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // HTTP Status - 200
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentBalance)
}
