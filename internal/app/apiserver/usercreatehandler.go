package apiserver

import (
	"AdHub/internal/app/models"
	"encoding/json"
	"net/http"
)

func (s *APIServer) UserCreateHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	newUser, err := s.Store.User().Create(&user)
	if err != nil {
		http.Error(w, "Error create user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newUser)
	w.Write(responseJSON)
}
