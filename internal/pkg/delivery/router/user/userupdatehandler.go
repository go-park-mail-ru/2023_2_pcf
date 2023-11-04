package router

import (
	"encoding/json"
	"io"
	"net/http"
)

func (ur *UserRouter) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		//Token       *string `json:"token"`
		Login       *string `json:"login"` // Любое поле может быть nil, если его нет в запросе
		Password    *string `json:"password"`
		FName       *string `json:"f_name"`
		LName       *string `json:"l_name"`
		CompanyName *string `json:"company_name"`
	}

	var avatar string
	err := r.ParseMultipartForm(10 << 20) // 10 MB - максимальный размер файла
	if err == nil {
		file, Header, err := r.FormFile("avatar")
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

	if err != nil {
		ur.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		ur.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

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
	uidAny := r.Context().Value("userid")
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
	if request.Login != nil {
		currentUser.Login = *request.Login
	}
	if request.Password != nil {
		currentUser.Password = *request.Password
	}
	if request.FName != nil {
		currentUser.FName = *request.FName
	}
	if request.LName != nil {
		currentUser.LName = *request.LName
	}
	if request.CompanyName != nil {
		currentUser.CompanyName = *request.CompanyName
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
