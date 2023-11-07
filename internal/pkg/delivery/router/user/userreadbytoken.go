package router

import (
	"encoding/json"
	"net/http"
)

func (ur *UserRouter) GetUserByTokenHandler(w http.ResponseWriter, r *http.Request) {

	uidAny := r.Context().Value("userid")
	userID, ok := uidAny.(int)
	if !ok {
		ur.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

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
