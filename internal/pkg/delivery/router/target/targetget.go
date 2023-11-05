package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (tr *TargetRouter) GetTargetHandler(w http.ResponseWriter, r *http.Request) {
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

	// Получение айди пользователя
	//userId, err := tr.Session.GetUserId(request.Token)
	//if err != nil {
	//	tr.logger.Error("Error getting user ID from session: " + err.Error())
	//	http.Error(w, "Unauthorized - Invalid session token", http.StatusUnauthorized)
	//	return
	//}

	uidAny := r.Context().Value("userid")
	userId, ok := uidAny.(int)
	if !ok {
		tr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
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
