package router

import (
	"encoding/json"
	"net/http"
)

func (mr *PadRouter) PadDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		PadID int `json:"pad_id"`
		//Token string `json:"token"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		mr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	uidAny := r.Context().Value("userId")
	userId, ok := uidAny.(int)
	if !ok {
		mr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	pad, err := mr.Pad.PadRead(request.PadID)
	if err != nil {
		mr.logger.Error("Error retrieving pad: " + err.Error())
		http.Error(w, "Error retrieving pad", http.StatusInternalServerError)
		return
	}

	// Проверка права пользователя
	if pad.Owner_id != userId {
		mr.logger.Error("User does not have permission to delete this pad")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Удаление
	err = mr.Pad.PadRemove(request.PadID)
	if err != nil {
		mr.logger.Error("Error deleting pad: " + err.Error())
		http.Error(w, "Error deleting pad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Pad deleted successfully"}`))
}
