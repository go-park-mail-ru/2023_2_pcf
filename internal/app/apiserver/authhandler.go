package apiserver

import (
	"AdHub/internal/app/models"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
)

type SessionHandler struct {
	Sessions map[string]int //key - session token, value - user_id
	mutex    sync.Mutex
}

func (sH *SessionHandler) AddSession(token string, userID int) {
	sH.mutex.Lock()
	defer sH.mutex.Unlock()

	sH.Sessions[token] = userID
}

type Session struct {
	token string `json:"token"`
}

var MySessionHandler = SessionHandler{
	Sessions: make(map[string]int),
}

func genToken(length int) (str string, err error) {
	b := make([]byte, length)
	_, err = rand.Read(b)
	if err != nil {
		return str, err
	}
	str = base64.URLEncoding.EncodeToString(b)
	return str, err
}

func GetUserFromRows(rows *sql.Rows, user *models.User) error {
	if rows.Next() {
		if err := rows.Scan(&user.Id, &user.Login, &user.Password, &user.FName, &user.LName); err != nil {
			return err
		}
	} else {
		return errors.New("user not found")
	}

	err := rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *APIServer) AuthHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	userSqlRows, err := s.Store.User().Get(user.Login)
	if err != nil {
		http.Error(w, "User check error: "+err.Error(), http.StatusBadRequest)
		return
	}

	userToCheck := &models.User{}
	err = GetUserFromRows(userSqlRows, userToCheck)
	if err != nil {
		//do smg...
	}

	if user.Password == userToCheck.Password {
		sToken, err := genToken(32)
		if err != nil {
			//do smg...
		}
		for _, exists := MySessionHandler.Sessions[sToken]; exists; _, exists = MySessionHandler.Sessions[sToken] {
			sToken, err = genToken(32)
			if err != nil {
				//do smg...
			}
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
	} else {
		http.Error(w, "User check error: "+err.Error(), http.StatusBadRequest)
		return
	}
}
