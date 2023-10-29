package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"net/http"
)

func (mr *AdRouter) AdListHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	id, err := mr.Session.GetUserId(token)
	if err != nil {
		mr.logger.Error("Session not found" + err.Error())
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	var ads []*entities.Ad
	ads, err = mr.Ad.AdReadList(id)
	if err != nil {
		mr.logger.Error("AdsList not found" + err.Error())
		http.Error(w, "Ads not found:", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ads); err != nil {
		mr.logger.Error("Bad request" + err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
