package router

import (
	"AdHub/survey/pkg/entities"
	"encoding/json"
	"net/http"
)

func (mr *SurveyRouter) GetSurvey(w http.ResponseWriter, r *http.Request) {
	var surveys []*entities.Survey
	surveys, err := mr.Survey.ReadList()
	if err != nil {
		mr.logger.Error(" not found" + err.Error())
		http.Error(w, " not found:", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(surveys); err != nil {
		mr.logger.Error("Bad request" + err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
