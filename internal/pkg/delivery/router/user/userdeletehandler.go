package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (mr *UserRouter) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userMail := vars["login"]

	if err := mr.User.UserDelete(userMail); err != nil {
		mr.logger.Error("Failed to delet user" + err.Error())
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // HTTP Status - 204
}
