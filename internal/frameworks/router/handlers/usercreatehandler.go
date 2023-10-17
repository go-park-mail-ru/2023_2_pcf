package apiserver

import (
	"AdHub/internal/app/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *APIServer) UserCreateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8081")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8081")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if !(user.ValidateEmail() && user.ValidatePassword() && user.ValidateFName() && user.ValidateLName()) {
		http.Error(w, "Invalid user params", http.StatusBadRequest)
		return
	}

	newUser, err := s.Store.User().Create(&user)
	if err != nil {
		http.Error(w, "Error create user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newUser)
	fmt.Println(responseJSON)
	w.Write(responseJSON)
}
