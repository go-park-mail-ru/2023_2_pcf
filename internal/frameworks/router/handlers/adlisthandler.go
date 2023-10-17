package apiserver

import (
	"AdHub/internal/app/auth"
	"AdHub/internal/app/models"
	"encoding/json"
	"log"
	"net/http"
)

func (s *APIServer) AdListHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	id, err := auth.MySessionStorage.GetUserId(token)
	if err != nil {
		log.Printf("Session not fount: %v", err)
		http.Error(w, "Session not fount", http.StatusNotFound)
		return
	}

	var ads []*models.Ad
	ads, err = s.Store.Ad().Read(id)
	if err != nil {
		log.Printf("Ads not found: %v", err)
		http.Error(w, "Ads not found:", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ads); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
