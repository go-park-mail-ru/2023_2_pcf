package router

import (
	"io"
	"net/http"
)

func (ur *UserRouter) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		//Token       *string `json:"token"`
		Login    string `json:"login"` // Любое поле может быть nil, если его нет в запросе
		Password string `json:"password"`
		FName    string `json:"f_name"`
		LName    string `json:"l_name"`
		Company  string `json:"s_name"`
	}

	var avatar string
	err := r.ParseMultipartForm(100 << 20) // 10 MB - максимальный размер файла
	if err == nil {
		file, Header, err := r.FormFile("avatar")
		if err == nil {
			defer file.Close()

			bfile, err := io.ReadAll(file)
			if err != nil {
				ur.logger.Error("Error reading file: " + err.Error())
				http.Error(w, "Error reading file", http.StatusInternalServerError)
				return
			}

			avatar, err = ur.File.Save(bfile, Header.Filename)
			if err != nil {
				ur.logger.Error("Error saving file: " + err.Error())
				http.Error(w, "Error saving file", http.StatusInternalServerError)
				return
			}
		}
	}

	err = r.ParseForm()
	if err != nil {
		ur.logger.Error("Error parsing form: " + err.Error())
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	request.Login = r.FormValue("login")
	request.Password = r.FormValue("password")
	request.FName = r.FormValue("f_name")
	request.LName = r.FormValue("l_name")
	request.Company = r.FormValue("s_name")

	// Проверяем, что ID пользователя предоставлен
	//if request.Token == nil {
	//	ur.logger.Error("Token is required")
	//	http.Error(w, "User ID is required", http.StatusBadRequest)
	//	return
	//}

	// Получение айди пользователя по сессии
	//userId, err := ur.Session.GetUserId(*request.Token)
	//if err != nil {
	//	ur.logger.Error("Error with session: " + err.Error())
	//	http.Error(w, "Error with authentication", http.StatusInternalServerError)
	//	return
	//}
	uidAny := r.Context().Value("userId")
	userId, ok := uidAny.(int)
	if !ok {
		ur.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	// Получение текущего состояния пользователя из бд
	currentUser, err := ur.User.UserReadById(userId)
	if err != nil {
		ur.logger.Error("Error retrieving user: " + err.Error())
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	// Обновляем поля, если они были предоставлены в запросе
	if request.Login != "" {
		currentUser.Login = request.Login
	}
	if request.Password != "" {
		currentUser.Password = request.Password
	}
	if request.FName != "" {
		currentUser.FName = request.FName
	}
	if request.LName != "" {
		currentUser.LName = request.LName
	}
	if request.Company != "" {
		currentUser.CompanyName = request.Company
	}
	if avatar != "" {
		currentUser.Avatar = avatar
	}

	// Обновление данных пользователя в базе данных
	err = ur.User.UserUpdate(currentUser)
	if err != nil {
		ur.logger.Error("Error updating user: ")
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
