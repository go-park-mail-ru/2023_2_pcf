package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"net/http"
)

func (tr *TargetRouter) CreateTargetHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token  string `json:"token"`
		Name   string `json:"name"`
		Gender string `json:"gender"`
		MinAge int    `json:"min_age"`
		MaxAge int    `json:"max_age"`
	}

	// Получение данных из запроса
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		tr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Получение айди пользователя из сессии
	ownerId, err := tr.Session.GetUserId(request.Token)
	if err != nil {
		tr.logger.Error("Error getting user ID from session: " + err.Error())
		http.Error(w, "Error getting user ID", http.StatusBadRequest)
		return
	}

	newTarget := entities.Target{
		Name:     request.Name,
		Owner_id: ownerId,
		Gender:   request.Gender,
		Min_age:  request.MinAge,
		Max_age:  request.MaxAge,
	}

	// Сохранение в бд
	targetCreated, err := tr.Target.TargetCreate(&newTarget)
	if err != nil {
		tr.logger.Error("Error creating new target: " + err.Error())
		http.Error(w, "Error creating new target", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(targetCreated); err != nil {
		tr.logger.Error("Error encoding response: " + err.Error())
	}
}
