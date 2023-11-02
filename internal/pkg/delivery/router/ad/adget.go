package router

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (mr *AdRouter) AdGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adIDStr, ok := vars["adID"]
	if !ok {
		mr.logger.Error("Ad ID is missing in the path")
		http.Error(w, "Ad ID is required", http.StatusBadRequest)
		return
	}
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

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ad)
}
