package router

import (
	"encoding/json"
	"net/http"
)

func (ar *AdRouter) AdGetAmountHandler(w http.ResponseWriter, r *http.Request) {

	var request struct {
		//Token  string `json:"token"`
		Amount int `json:"amount"`
		Offset int `json:"offset"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		ar.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//Проверка входных данных
	if request.Offset < 0 || request.Amount < 0 {
		ar.logger.Error("Incorrect request data. Amount and offset must be >= 0")
		http.Error(w, "Incorrect request data. Amount and offset must be >= 0", http.StatusBadRequest)
		return
	}

	// Получение ID пользователя через сесссию
	//userId, err := ar.Session.GetUserId(request.Token)
	//if err != nil {
	//	ar.logger.Error("Error getting user ID from session: " + err.Error())
	//	http.Error(w, "Unauthorized - Invalid session token", http.StatusUnauthorized)
	//	return
	//}
	uidAny := r.Context().Value("userid")
	userId, ok := uidAny.(int)
	if !ok {
		ar.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	// Получение списка объявлений
	ads, err := ar.Ad.AdReadList(userId)
	if err != nil {
		ar.logger.Error("Error reading ads list: " + err.Error())
		http.Error(w, "Error reading ads list", http.StatusInternalServerError)
		return
	}

	if request.Offset > len(ads) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ads[0:0]); err != nil {
			ar.logger.Error("Error encoding response: " + err.Error())
		}
	}

	if request.Amount == 0 || request.Offset+request.Amount > len(ads) {
		request.Amount = len(ads)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ads[request.Offset:request.Amount]); err != nil {
		ar.logger.Error("Error encoding response: " + err.Error())
	}

}
