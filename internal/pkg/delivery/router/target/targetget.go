package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (tr *TargetRouter) GetTargetHandler(w http.ResponseWriter, r *http.Request) {
	// Получение данных из запроса
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		tr.logger.Error("No target id provided")
		http.Error(w, "No target id provided", http.StatusBadRequest)
		return
	}

	targetId, err := strconv.Atoi(id)
	if err != nil {
		tr.logger.Error("Invalid target id: " + err.Error())
		http.Error(w, "Invalid target id", http.StatusBadRequest)
		return
	}

	var request struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		tr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Получение айди пользователя
	userId, err := tr.Session.GetUserId(request.Token)
	if err != nil {
		tr.logger.Error("Error getting user ID from session: " + err.Error())
		http.Error(w, "Unauthorized - Invalid session token", http.StatusUnauthorized)
		return
	}

	// Получение таргета из бд
	target, err := tr.Target.TargetRead(targetId)
	if err != nil {
		tr.logger.Error("Error retrieving target: " + err.Error())
		http.Error(w, "Target not found", http.StatusNotFound)
		return
	}

	// Проверка прав пользователя
	if target.Owner_id != userId {
		tr.logger.Error("Access denied - User does not own the requested target")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(target); err != nil {
		tr.logger.Error("Error encoding response: " + err.Error())
	}
}
