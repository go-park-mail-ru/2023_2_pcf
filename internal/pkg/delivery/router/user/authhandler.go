package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"net/http"
	"time"
)

func (mr *UserRouter) AuthHandler(w http.ResponseWriter, r *http.Request) {
	//Парсинг юзера из запроса
	var user entities.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		mr.logger.Error("Invalid request body.")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//Валидация пришедших данных
	if !(user.ValidateEmail() && user.ValidatePassword()) {
		mr.logger.Error("Invalid user params.")
		http.Error(w, "Invalid user params", http.StatusBadRequest)
		return
	}

	//Получение реального юзера из бд по логину из запроса
	userFromDB, err := mr.User.UserReadByLogin(user.Login)
	if err != nil {
		mr.logger.Error("Error while getting user from DB: " + err.Error())
		http.Error(w, "Error while getting user from DB: "+err.Error(), http.StatusBadRequest)
		return
	}

	//Проверка пароля
	if user.Password == userFromDB.Password {
		newSession, err := mr.Session.Auth(userFromDB)
		if err != nil {
			mr.logger.Error("Error while auth token generation" + err.Error())
			http.Error(w, "Error while token gen", http.StatusInternalServerError)
		}

		newCsrf, err := mr.Csrf.CsrfCreate(userFromDB.Id)
		if err != nil {
			mr.logger.Error("Error while csrf token generation" + err.Error())
			http.Error(w, "Error while csrf token gen", http.StatusInternalServerError)
		}

		//Кукисет и возврат ответа (успех)
		cookie := &http.Cookie{
			Name:     "session_token",
			Value:    newSession.Token,
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
			Domain:   "127.0.0.1",
			Path:     "/",
		}
		http.SetCookie(w, cookie)

		cookie2 := &http.Cookie{
			Name:     "csrf_token",
			Value:    newCsrf.Token,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Domain:   "127.0.0.1",
			Path:     "/",
		}
		http.SetCookie(w, cookie2)

		w.WriteHeader(http.StatusOK)

	} else {
		http.Error(w, "Wrong password", http.StatusBadRequest)
		return
	}
}
