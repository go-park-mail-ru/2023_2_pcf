package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (tr *TargetRouter) UpdateTargetHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Gender    string `json:"gender"`
		MinAge    string `json:"min_age"`
		MaxAge    string `json:"max_age"`
		Interests string `json:"interests"`
		Tags      string `json:"tags"`
		Keys      string `json:"keys"`
		Regions   string `json:"regions"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		tr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Получение айди пользователя из сессии
	uidAny := r.Context().Value("userId")
	userId, ok := uidAny.(int)
	if !ok {
		tr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

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
		Owner_id:  userId,
		Gender:    request.Gender,
		Min_age:   min,
		Max_age:   max,
		Interests: interests,
		Tags:      tags,
		Keys:      keys,
		Regions:   regions,
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
	if request.Name != "" {
		currentTarget.Name = newTarget.Name
	}
	if request.Gender != "" {
		currentTarget.Gender = newTarget.Gender
	}
	if request.MinAge != "" {
		currentTarget.Min_age = newTarget.Min_age
	}
	if request.MaxAge != "" {
		currentTarget.Max_age = newTarget.Max_age
	}
	if request.Interests != "" {
		currentTarget.Interests = strings.Split(request.Interests, ", ")
	}
	if request.Tags != "" {
		currentTarget.Tags = strings.Split(request.Tags, ", ")
	}
	if request.Keys != "" {
		currentTarget.Keys = strings.Split(request.Keys, ", ")
	}
	if request.Regions != "" {
		currentTarget.Regions = strings.Split(request.Regions, ", ")
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
