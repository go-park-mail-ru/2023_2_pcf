package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"fmt"
	"net/http"
)

func (mr *UserRouter) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		mr.logger.Error("Invalid request body." + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if !(user.ValidateEmail() && user.ValidatePassword() && user.ValidateFName() && user.ValidateLName()) {
		mr.logger.Error("Invalid user params.")
		http.Error(w, "Invalid user params", http.StatusBadRequest)
		return
	}

	newUser, err := mr.User.UserCreate(&user)
	if err != nil {
		mr.logger.Error("Error user create" + err.Error())
		http.Error(w, "Error create user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newUser)
	fmt.Println(responseJSON)
	w.Write(responseJSON)
}
