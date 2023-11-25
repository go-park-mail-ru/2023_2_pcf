package router

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func (mr *SurveyRouter) GetStat(w http.ResponseWriter, r *http.Request) {
	SurveyIdStr := r.URL.Query().Get("id")
	SID, err := strconv.Atoi(SurveyIdStr)
	if err != nil {
		mr.logger.Error("Invalid Survey ID: " + err.Error())
		http.Error(w, "Invalid Survey ID", http.StatusBadRequest)
		return
	}

	// Получаем объявление из бд
	s, err := mr.Rate.RateRead(SID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "rate not found", http.StatusNotFound)
		} else {
			mr.logger.Error("Error retrieving rate: " + err.Error())
			http.Error(w, "Error retrieving rate", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}
