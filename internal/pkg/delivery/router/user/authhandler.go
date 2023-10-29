package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/auth"
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
	userFromDB, err := mr.User.UserRead(user.Login)
	if err != nil {
		mr.logger.Error("Error while getting user from DB: " + err.Error())
		http.Error(w, "Error while getting user from DB: "+err.Error(), http.StatusBadRequest)
		return
	}

	//Проверка пароля
	if user.Password == userFromDB.Password {
		newSession := auth.Session{UserId: userFromDB.Id}
		err = newSession.SetToken()
		if err != nil {
			mr.logger.Error("Error while token generation" + err.Error())
			http.Error(w, "Error while token gen", http.StatusInternalServerError)
			return
		}

		//Проверка уникальности токена, перегенерация если он уже занят
		for contains := auth.MySessionStorage.Contains(newSession.Token); contains; auth.MySessionStorage.Contains(newSession.Token) {
			err = newSession.SetToken()
			if err != nil {
				mr.logger.Error("Error while token generation" + err.Error())
				http.Error(w, "Error while token gen", http.StatusInternalServerError)
				return
			}
		}
		auth.MySessionStorage.AddSession(newSession)

		//Перевод структуры сессии в JSON
		w.Header().Set("Content-Type", "application/json")
		responseJSON, err := json.Marshal(newSession)
		if err != nil {
			defer auth.MySessionStorage.RemoveSession(newSession.Token)
			mr.logger.Error("Failed to marshal JSON." + err.Error())
			http.Error(w, "Failed to marshal JSON:", http.StatusInternalServerError)
			return
		}

		//Кукисет и возврат ответа (успех)
		cookie := &http.Cookie{
			Name:     "session_token",
			Value:    newSession.Token,
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)

	} else {
		http.Error(w, "Wrong password", http.StatusBadRequest)
		return
	}
}
