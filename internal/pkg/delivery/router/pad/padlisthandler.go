package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"net/http"
)

func (mr *PadRouter) PadListHandler(w http.ResponseWriter, r *http.Request) {
	uidAny := r.Context().Value("userId")
	id, ok := uidAny.(int)
	if !ok {
		mr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	var pads []*entities.Pad
	pads, err := mr.Pad.PadReadList(id)
	if err != nil {
		mr.logger.Error("PadsList not found" + err.Error())
		http.Error(w, "Pads not found:", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pads); err != nil {
		mr.logger.Error("Bad request" + err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
