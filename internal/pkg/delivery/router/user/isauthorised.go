package router

import (
	"AdHub/proto/api"
	"context"
	"encoding/json"
	"net/http"
)

func (ur *UserRouter) IsAuthorisedHandler(w http.ResponseWriter, r *http.Request) {

	isAuth := "false"

	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		isAuth = "false"
	} else {
		_, err = ur.Session.GetUserId(context.Background(), &api.GetRequest{Token: sessionToken.Value})
		if err != nil {
			isAuth = "false"
		} else {
			isAuth = "true"
		}
	}

	response := struct {
		IsAuth string `json:"is_auth"`
	}{
		IsAuth: isAuth,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
