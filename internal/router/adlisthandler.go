package router

import (
	"AdHub/internal/entities"
	"AdHub/pkg/auth"
	"encoding/json"
	"log"
	"net/http"
)

func (mr *MuxRouter) AdListHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	id, err := auth.MySessionStorage.GetUserId(token)
	if err != nil {
		log.Printf("Session not fount: %v", err)
		http.Error(w, "Session not fount", http.StatusNotFound)
		return
	}

	var ads []*entities.Ad
	ads, err = mr.Ad.AdGetList(id)
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
