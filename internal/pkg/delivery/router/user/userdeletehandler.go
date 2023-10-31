package router

import (
	"net/http"
)

func (mr *UserRouter) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")

	if err := mr.User.UserDelete(login); err != nil {
		mr.logger.Error("Failed to delete user" + err.Error())
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // HTTP Status - 204
}
