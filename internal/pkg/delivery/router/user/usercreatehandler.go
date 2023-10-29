package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (mr *UserRouter) UserCreateHandler(w http.ResponseWriter, r *http.Request) {

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

	//Получение данных из запроса
	var user entities.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		mr.logger.Error("Invalid request body." + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//Валидация полученных данных
	if !(user.ValidateEmail() && user.ValidatePassword() && user.ValidateFName() && user.ValidateLName()) {
		mr.logger.Error("Invalid user params.")
		http.Error(w, "Invalid user params", http.StatusBadRequest)
		return
	}

	//Создание юзера
	newUser, err := mr.User.UserCreate(&user)
	if err != nil {
		mr.logger.Error("Error user create" + err.Error())

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			http.Error(w, "User with this login already exists", http.StatusConflict)
		} else {
			http.Error(w, "Error create user", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newUser)
	fmt.Println(responseJSON)
	w.Write(responseJSON)
}
