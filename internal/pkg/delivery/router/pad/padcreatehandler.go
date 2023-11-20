package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"net/http"
	"strconv"
)

func (mr *PadRouter) PadCreateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		WebsiteLink string `json:"website_link"`
		Price       string `json:"price"`
		TargetId    string `json:"target_id"`
	}

	request.Name = r.FormValue("name")
	request.Description = r.FormValue("description")
	request.WebsiteLink = r.FormValue("website_link")
	request.Price = r.FormValue("price")
	request.TargetId = r.FormValue("target_id")

	uidAny := r.Context().Value("userId")
	uid, ok := uidAny.(int)
	if !ok {
		mr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}
	newbudget, err := strconv.ParseFloat(request.Price, 64)
	if err != nil {
		mr.logger.Error("Error budget parse" + err.Error())
		http.Error(w, "Error budget parse", http.StatusInternalServerError)
		return
	}

	target, err := strconv.Atoi(request.TargetId)
	if err != nil {
		mr.logger.Error("Error target parse" + err.Error())
		http.Error(w, "Error target parse", http.StatusInternalServerError)
		return
	}

	pad := entities.Pad{
		Name:         request.Name,
		Description:  request.Description,
		Website_link: request.WebsiteLink,
		Price:        newbudget, // Преобразование int в float64
		Owner_id:     uid,       // Укажите нужное значение Owner_id
		Target_id:    target,    // Укажите нужное значение Target_id
	}

	newPad, err := mr.Pad.PadCreate(&pad)
	if err != nil {
		mr.logger.Error("Error ad create" + err.Error())
		http.Error(w, "Error create ad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newPad)
	w.Write(responseJSON)
}
