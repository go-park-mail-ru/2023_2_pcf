package router

import (
	"AdHub/survey/pkg/entities"
	"encoding/json"
	"net/http"
	"strconv"
)

func (mr *SurveyRouter) RateCreateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		User_id   string `json:"user_id"`
		Rate      string `json:"rate"`
		Survey_id string `json:"survey_id"`
	}

	err := r.ParseForm()
	if err != nil {
		mr.logger.Error("Error parsing form: " + err.Error())
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	request.User_id = r.FormValue("user_id")
	request.Rate = r.FormValue("rate")
	request.Survey_id = r.FormValue("survey_id")

	newuid, err := strconv.Atoi(request.User_id)
	if err != nil {
		mr.logger.Error("Error user id parse" + err.Error())
		http.Error(w, "Error user id parse", http.StatusInternalServerError)
		return
	}

	newrate, err := strconv.Atoi(request.Rate)
	if err != nil {
		mr.logger.Error("Error rate parse" + err.Error())
		http.Error(w, "Error rate parse", http.StatusInternalServerError)
		return
	}

	newsid, err := strconv.Atoi(request.Survey_id)
	if err != nil {
		mr.logger.Error("Error survey parse" + err.Error())
		http.Error(w, "Error survey parse", http.StatusInternalServerError)
		return
	}

	rate := entities.Rate{
		Id:        0,
		User_id:   newuid,
		Rate:      newrate,
		Survey_id: newsid,
	}

	newRate, err := mr.Rate.RateCreate(&rate)
	if err != nil {
		mr.logger.Error("Error ad create" + err.Error())
		http.Error(w, "Error create ad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newRate)
	w.Write(responseJSON)
}
