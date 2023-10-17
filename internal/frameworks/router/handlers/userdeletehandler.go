package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userMail := vars["login"]

	if err := s.Store.User().Remove(userMail); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // HTTP Status - 204
}
