package router

import (
	"encoding/json"
	"net/http"
)

func (tr *TargetRouter) TargetDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Id    int    `json:"id"`
		Token string `json:"token"`
	}

	// Получение данных из запроса
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		tr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Получение ID пользователя из сессии
	userId, err := tr.Session.GetUserId(request.Token)
	if err != nil {
		tr.logger.Error("Error getting user ID from session: " + err.Error())
		http.Error(w, "Error getting user ID", http.StatusUnauthorized)
		return
	}

	// Проверка прав пользователя
	target, err := tr.Target.TargetRead(request.Id)
	if err != nil {
		tr.logger.Error("Error retrieving target: " + err.Error())
		http.Error(w, "Target not found", http.StatusNotFound)
		return
	}
	if target.Owner_id != userId {
		tr.logger.Error("User does not have permission to delete this target")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Удаление таргета
	if err := tr.Target.TargetRemove(request.Id); err != nil {
		tr.logger.Error("Error deleting target: " + err.Error())
		http.Error(w, "Error deleting target", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
