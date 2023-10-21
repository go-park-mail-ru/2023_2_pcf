package router

import (
	"encoding/json"
	"net/http"
)

func (mr *UserRouter) UserReadHandler(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")

	user, err := mr.User.UserRead(login)
	if err != nil {
		mr.logger.Error("User not found" + err.Error())
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		mr.logger.Error("Bad request" + err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
