package apiserver

import (
	"AdHub/internal/app/auth"
	"AdHub/internal/app/models"
	"encoding/json"
	"net/http"
	"time"
)

func (s *APIServer) AuthHandler(w http.ResponseWriter, r *http.Request) {

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

		w.WriteHeader(http.StatusCreated) // HTTP Status - 201
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
		w.Write(responseJSON)

	} else {
		http.Error(w, "User check error: "+err.Error(), http.StatusBadRequest)
		return
	}
}
