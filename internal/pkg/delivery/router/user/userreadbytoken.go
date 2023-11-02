package router

import (
	"encoding/json"
	"net/http"
)

func (ur *UserRouter) GetUserByTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token string `json:"token"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		ur.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Получение ID пользователя по токену
	userID, err := ur.Session.GetUserId(request.Token)
	if err != nil {
		ur.logger.Error("Error getting user ID from session: " + err.Error())
		http.Error(w, "Error getting session", http.StatusBadRequest)
		return
	}

	// Получение данных пользователя
	user, err := ur.User.UserReadById(userID)
	if err != nil {
		ur.logger.Error("Error getting user by ID: " + err.Error())
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		ur.logger.Error("Error encoding response: " + err.Error())
		http.Error(w, "Error encoding user data", http.StatusInternalServerError)
	}
}
