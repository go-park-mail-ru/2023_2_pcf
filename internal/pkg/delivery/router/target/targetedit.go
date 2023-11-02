package router

import (
	"encoding/json"
	"net/http"
)

func (tr *TargetRouter) UpdateTargetHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token  string  `json:"token"`
		Id     int     `json:"id"`
		Name   *string `json:"name,omitempty"`
		Gender *string `json:"gender,omitempty"`
		MinAge *int    `json:"min_age,omitempty"`
		MaxAge *int    `json:"max_age,omitempty"`
	}

	// Получение данных из запроса
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		tr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Получение айди пользователя из сессии
	userId, err := tr.Session.GetUserId(request.Token)
	if err != nil {
		tr.logger.Error("Error getting user ID from session: " + err.Error())
		http.Error(w, "Error getting user ID", http.StatusBadRequest)
		return
	}

	// Получение текущего таргета из бд
	currentTarget, err := tr.Target.TargetRead(request.Id)
	if err != nil {
		tr.logger.Error("Error retrieving target: " + err.Error())
		http.Error(w, "Error retrieving target", http.StatusInternalServerError)
		return
	}

	// Проверка прав владельца
	if currentTarget.Owner_id != userId {
		tr.logger.Error("User does not have permission to update this target")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Обновление полей
	if request.Name != nil {
		currentTarget.Name = *request.Name
	}
	if request.Gender != nil {
		currentTarget.Gender = *request.Gender
	}
	if request.MinAge != nil {
		currentTarget.Min_age = *request.MinAge
	}
	if request.MaxAge != nil {
		currentTarget.Max_age = *request.MaxAge
	}

	// Сохранение в бд
	if err := tr.Target.TargetUpdate(currentTarget); err != nil {
		tr.logger.Error("Error updating target: " + err.Error())
		http.Error(w, "Error updating target", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(currentTarget); err != nil {
		tr.logger.Error("Error encoding response: " + err.Error())
	}
}
