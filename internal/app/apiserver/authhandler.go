package apiserver

import (
	"AdHub/internal/app/auth"
	"AdHub/internal/app/models"
	"encoding/json"
	"net/http"
	"time"
)

func (s *APIServer) AuthHandler(w http.ResponseWriter, r *http.Request) {

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

	userFromDB, err := s.Store.User().Read(user.Login)
	if err != nil {
		http.Error(w, "Error while getting user from DB: "+err.Error(), http.StatusBadRequest)
		return
	}

	if user.Password == userFromDB.Password {
		newSession := auth.Session{UserId: userFromDB.Id}
		err = newSession.SetToken()
		if err != nil {
			//do smg...
		}
		//check if this token already exists
		for contains := auth.MySessionStorage.Contains(newSession.Token); contains; auth.MySessionStorage.Contains(newSession.Token) {
			err = newSession.SetToken()
			if err != nil {
				//do smg...
			}
		}
		auth.MySessionStorage.AddSession(newSession)

		w.Header().Set("Content-Type", "application/json")
		responseJSON, err := json.Marshal(newSession)
		if err != nil {
			defer auth.MySessionStorage.RemoveSession(newSession.Token)
			http.Error(w, "Failed to marshal JSON: "+err.Error(), http.StatusInternalServerError)
			return
		}

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