package router

import (
	"net/http"
)

func (mr *UserRouter) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {

	uidAny := r.Context().Value("userid")
	userId, ok := uidAny.(int)
	if !ok {
		mr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	login := r.URL.Query().Get("login")
	userFromDb, err := mr.User.UserReadByLogin(login)
	if err != nil {
		mr.logger.Error("Error while getting user from db" + err.Error())
		http.Error(w, "Error while authentication", http.StatusInternalServerError)
		return
	}

	if userId != userFromDb.Id {
		mr.logger.Error("Access denied" + err.Error())
		http.Error(w, "Access denied", http.StatusMethodNotAllowed)
		return
	}

	if err := mr.User.UserDelete(login); err != nil {
		mr.logger.Error("Failed to delete user" + err.Error())
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // HTTP Status - 204
}
