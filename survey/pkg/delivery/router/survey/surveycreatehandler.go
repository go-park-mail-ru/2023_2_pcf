package router

import (
	"AdHub/survey/pkg/entities"
	"encoding/json"
	"net/http"
	"strconv"
)

func (mr *SurveyRouter) SurveyCreateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Question string `json:"question"`
		Type     string `json:"type"`
	}

	err := r.ParseForm()
	if err != nil {
		mr.logger.Error("Error parsing form: " + err.Error())
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	request.Question = r.FormValue("question")
	request.Type = r.FormValue("type")

	var newtype int
	newtype, err = strconv.Atoi(request.Type)
	if err != nil {
		mr.logger.Error("Error type parse" + err.Error())
		http.Error(w, "Error type parse", http.StatusInternalServerError)
		return
	}

	survey := entities.Survey{
		Id:       0,
		Question: request.Question,
		Type:     newtype,
	}

	newSurvey, err := mr.Survey.SurveyCreate(&survey)
	if err != nil {
		mr.logger.Error("Error ad create" + err.Error())
		http.Error(w, "Error create ad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newSurvey)
	w.Write(responseJSON)
}
