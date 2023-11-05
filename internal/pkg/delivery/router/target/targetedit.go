package router

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (tr *TargetRouter) UpdateTargetHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Id        int     `json:"id"`
		Name      *string `json:"name,omitempty"`
		Gender    *string `json:"gender,omitempty"`
		MinAge    *int    `json:"min_age,omitempty"`
		MaxAge    *int    `json:"max_age,omitempty"`
		Interests string  `json:"interests,omitempty"`
		Tags      string  `json:"tags,omitempty"`
		Keys      string  `json:"keys,omitempty"`
		Regions   string  `json:"regions,omitempty"`
	}

	// Получение данных из запроса
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		tr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Получение айди пользователя из сессии
	uidAny := r.Context().Value("userid")
	userId, ok := uidAny.(int)
	if !ok {
		tr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
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
