package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *APIServer) UserReadHandler(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")

	user, err := s.Store.User().Read(login)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
