package router

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func (mr *AdRouter) AdGetHandler(w http.ResponseWriter, r *http.Request) {
	adIDStr := r.URL.Query().Get("adID")
	adID, err := strconv.Atoi(adIDStr)
	if err != nil {
		mr.logger.Error("Invalid Ad ID: " + err.Error())
		http.Error(w, "Invalid Ad ID", http.StatusBadRequest)
		return
	}

	// Получаем объявление из бд
	ad, err := mr.Ad.AdRead(adID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Ad not found", http.StatusNotFound)
		} else {
			mr.logger.Error("Error retrieving ad: " + err.Error())
			http.Error(w, "Error retrieving ad", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ad)
}
