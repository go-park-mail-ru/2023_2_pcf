package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (tr *TargetRouter) CreateTargetHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name      string `json:"name"`
		Gender    string `json:"gender"`
		MinAge    string `json:"min_age"`
		MaxAge    string `json:"max_age"`
		Interests string `json:"interests"`
		Tags      string `json:"tags"`
		Keys      string `json:"keys"`
		Regions   string `json:"regions"`
	}

	// Получение данных из запроса
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		tr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Получение айди пользователя из сессии
	uidAny := r.Context().Value("userId")
	ownerID, ok := uidAny.(int)
	if !ok {
		tr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	// Разделение тегов, интересов, ключей и регионов
	interests := strings.Split(request.Interests, ", ")
	tags := strings.Split(request.Tags, ", ")
	keys := strings.Split(request.Keys, ", ")
	regions := strings.Split(request.Regions, ", ")
	min, err := strconv.Atoi(request.MinAge)
	if err != nil {
		tr.logger.Error("Invalid min age: " + err.Error())
		http.Error(w, "Invalid min age", http.StatusBadRequest)
		return
	}

	max, err := strconv.Atoi(request.MaxAge)
	if err != nil {
		tr.logger.Error("Invalid max age: " + err.Error())
		http.Error(w, "Invalid max age", http.StatusBadRequest)
		return
	}

	newTarget := entities.Target{
		Name:      request.Name,
		Owner_id:  ownerID,
		Gender:    request.Gender,
		Min_age:   min,
		Max_age:   max,
		Interests: interests,
		Tags:      tags,
		Keys:      keys,
		Regions:   regions,
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
