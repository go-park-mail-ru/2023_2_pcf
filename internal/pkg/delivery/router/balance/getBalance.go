package router

import (
	"encoding/json"
	"net/http"
)

func (br *BalanceRouter) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {

	//if request.Token == "" {
	//	br.logger.Error("Token is required")
	//	http.Error(w, "Token is required", http.StatusBadRequest)
	//	return
	//}

	//userId, err := br.Session.GetUserId(request.Token)
	//if err != nil {
	//	br.logger.Error("Error getting user ID from session: " + err.Error())
	//	http.Error(w, "Error getting user ID", http.StatusBadRequest)
	//	return
	//}
	uidAny := r.Context().Value("userId")
	userId, ok := uidAny.(int)
	if !ok {
		br.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	// Получение баланса пользователя
	balance, err := br.Balance.BalanceRead(userId)
	if err != nil {
		br.logger.Error("Error retrieving balance: " + err.Error())
		http.Error(w, "Error retrieving balance", http.StatusInternalServerError)
		return
	}

	// Отправка данных о балансе пользователю
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(balance); err != nil {
		br.logger.Error("Error encoding response: " + err.Error())
	}
}
