package router

import (
	"net/http"
	"strconv"
)

func (mr *PadRouter) PadUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		PadId       string `json:"id"`
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
	request.PadId = r.FormValue("id")
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
	id, err := strconv.Atoi(request.PadId)
	if err != nil {
		mr.logger.Error("Error ad parse" + err.Error())
		http.Error(w, "Error ad parse", http.StatusInternalServerError)
		return
	}

	// Получение текущего состояния рекламы из базы данных
	currentPad, err := mr.Pad.PadRead(id)
	if err != nil {
		mr.logger.Error("Error retrieving pad: " + err.Error())
		http.Error(w, "Error retrieving pad", http.StatusInternalServerError)
		return
	}

	// Проверяем, что реклама принадлежит пользователю из сессии
	if currentPad.Owner_id != uid {
		mr.logger.Error("User does not have permission to update this pad")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Обновляем поля рекламы
	if request.Name != "" {
		currentPad.Name = request.Name
	}
	if request.Description != "" {
		currentPad.Description = request.Description
	}
	if request.WebsiteLink != "" {
		currentPad.Website_link = request.WebsiteLink
	}
	if request.Price != "" {
		currentPad.Price = newbudget
	}
	if request.TargetId != "" {
		currentPad.Target_id = target
	}

	// Обновление данных рекламы в базе данных
	err = mr.Pad.PadUpdate(currentPad)
	if err != nil {
		mr.logger.Error("Error updating pad: " + err.Error())
		http.Error(w, "Error updating pad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
