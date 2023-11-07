package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"fmt"
	"net/http"
)

func (tr *TargetRouter) TargetListHandler(w http.ResponseWriter, r *http.Request) {
	uidAny := r.Context().Value("userId")
	userId, ok := uidAny.(int)
	if !ok {
		tr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	var target []*entities.Target

	target, err := tr.Target.TargetReadList(userId)
	if err != nil {
		tr.logger.Error("Error retrieving target: " + err.Error())
		http.Error(w, "Target not found", http.StatusNotFound)
		return
	}
	if len(target) != 0 {
		fmt.Println(target)
		if target[0].Owner_id != userId {
			tr.logger.Error("Access denied - User does not own the requested target")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(target); err != nil {
		tr.logger.Error("Error encoding response: " + err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
