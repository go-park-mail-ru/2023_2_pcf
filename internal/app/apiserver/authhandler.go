package apiserver

import (
	"AdHub/internal/app/models"
	"encoding/json"
	"net/http"
	"crypto/rand"
	"sync"
)

type SessionHandler struct {
	Sessions map[string]int //key - session token, value - user_id
	mutex sync.Mutex
}

func (sH *sessionHandler) AddSession(token string, userID int) {
	sH.mutex.Lock()
	defer sH.mutex.Unlock()

	sH.sessions[token] = userID
}

type Session struct {
	token string `json:"token"`
}

MySessionHandler := sessionHandler{
	Sessions: make(map[string]int),
}

func genToken(lenght int) string, err{
	b := make([]byte, length)
	_, err := rand.Read(b)
	if(err != nil){
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {

	var user User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	userToCheck, err := s.Store.User().Get(user.Login)
	if(err != nil){
		http.Error(w, "User check error: " + err.Error(), http.StatusBadRequest)
		return
	}

	if(user.Password == userToCheck.Password){
		sToken, err := genToken(32)
		if(err != nil){
			//do smg...
		}
		for _, exists := MySessionHandler.Sessions[sToken]; exists; _, exists = MySessionHandler.Sessions[sToken]{
			sToken := genToken(32)
		}
		MySessionHandler.AddSession(sToken, userToCheck.Id)

		newSession := Session{token: sToken}
		w.WriteHeader(http.StatusCreated) // HTTP Status - 201
		w.Header().Set("Content-Type", "application/json")
		responseJSON, err := json.Marshal(newSession)
		if err != nil {
			http.Error(w, "Failed to marshal JSON: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(responseJSON)
		return
	} else{
		http.Error(w, "User check error: " + err.Error(), http.StatusBadRequest)
		return
	}
}
