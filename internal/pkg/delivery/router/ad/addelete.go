package router

import (
	"encoding/json"
	"net/http"
)

func (mr *AdRouter) AdDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AdID  int    `json:"ad_id"`
		Token string `json:"token"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		mr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Получение ID пользователя из сессии
	userId, err := mr.Session.GetUserId(request.Token)
	if err != nil {
		mr.logger.Error("Error getting session: " + err.Error())
		http.Error(w, "Error getting session", http.StatusBadRequest)
		return
	}

	// Получение объявления из бд
	ad, err := mr.Ad.AdRead(request.AdID)
	if err != nil {
		mr.logger.Error("Error retrieving ad: " + err.Error())
		http.Error(w, "Error retrieving ad", http.StatusInternalServerError)
		return
	}

	// Проверка права пользователя
	if ad.Owner_id != userId {
		mr.logger.Error("User does not have permission to delete this ad")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Удаление
	err = mr.Ad.AdRemove(request.AdID)
	if err != nil {
		mr.logger.Error("Error deleting ad: " + err.Error())
		http.Error(w, "Error deleting ad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Ad deleted successfully"}`))
}
