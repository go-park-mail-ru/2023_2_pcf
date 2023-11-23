package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/proto/api"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func (mr *UserRouter) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
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

	newBalance := &entities.Balance{
		Total_balance:     0,
		Reserved_balance:  0,
		Available_balance: 0,
	}

	balance, err := mr.Balance.BalanceCreate(newBalance)
	if err != nil {
		mr.logger.Error("Invalid balance create.")
		http.Error(w, "Invalid balance create.", http.StatusBadRequest)
		return
	}

	newUser.BalanceId = balance.Id
	err = mr.User.UserUpdate(newUser)
	if err != nil {
		mr.logger.Error("Invalid user update.")
		http.Error(w, "Invalid user update.", http.StatusBadRequest)
		return
	}

	newSession, err := mr.Session.Auth(context.Background(), &api.AuthRequest{Id: int64(newUser.Id), Login: newUser.Login, Password: newUser.Password, FName: newUser.FName, LName: newUser.LName, CompanyName: newUser.CompanyName, Avatar: newUser.Avatar, BalanceId: int64(newUser.BalanceId)})
	if err != nil {
		mr.logger.Error("Error while token generation" + err.Error())
		http.Error(w, "Error while token gen", http.StatusInternalServerError)
	}

	//Кукисет и возврат ответа (успех)
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    newSession.GetToken(),
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Domain:   "127.0.0.1",
		Path:     "/",
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newUser)
	fmt.Println(responseJSON)
	w.Write(responseJSON)
}
