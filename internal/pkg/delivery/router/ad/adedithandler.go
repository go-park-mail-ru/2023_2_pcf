package router

import (
	"encoding/json"
	"io"
	"net/http"
)

func (mr *AdRouter) AdUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		//Token       string   `json:"token"`
		AdId        int      `json:"ad_id"`
		Name        *string  `json:"name"`
		Description *string  `json:"description"`
		WebsiteLink *string  `json:"website_link"`
		Budget      *float64 `json:"budget"`
		TargetId    *int     `json:"target_id"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		mr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var image string
	err := r.ParseMultipartForm(10 << 20) // 10 MB - максимальный размер файла
	if err == nil {
		file, Header, err := r.FormFile("image")
		defer file.Close()

		bfile, err := io.ReadAll(file)
		if err != nil {
			mr.logger.Error("Error reading file: " + err.Error())
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		image, err = mr.File.Save(bfile, Header.Filename)
		if err != nil {
			mr.logger.Error("Error saving file: " + err.Error())
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}
	}

	// Получаем ID пользователя из сессии
	//userId, err := mr.Session.GetUserId(request.Token)
	//if err != nil {
	//	mr.logger.Error("Error getting session: " + err.Error())
	//	http.Error(w, "Error getting session", http.StatusBadRequest)
	//	return
	//}
	uidAny := r.Context().Value("userid")
	userId, ok := uidAny.(int)
	if !ok {
		mr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	// Получение текущего состояния рекламы из базы данных
	currentAd, err := mr.Ad.AdRead(request.AdId)
	if err != nil {
		mr.logger.Error("Error retrieving ad: " + err.Error())
		http.Error(w, "Error retrieving ad", http.StatusInternalServerError)
		return
	}

	// Проверяем, что реклама принадлежит пользователю из сессии
	if currentAd.Owner_id != userId {
		mr.logger.Error("User does not have permission to update this ad")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Обновляем поля рекламы
	if request.Name != nil {
		currentAd.Name = *request.Name
	}
	if request.Description != nil {
		currentAd.Description = *request.Description
	}
	if request.WebsiteLink != nil {
		currentAd.Website_link = *request.WebsiteLink
	}
	if request.Budget != nil {
		currentAd.Budget = *request.Budget
	}
	if request.TargetId != nil {
		currentAd.Target_id = *request.TargetId
	}
	if image != "" {
		currentAd.Image_link = image
	}

	// Обновление данных рекламы в базе данных
	err = mr.Ad.AdUpdate(currentAd)
	if err != nil {
		mr.logger.Error("Error updating ad: " + err.Error())
		http.Error(w, "Error updating ad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
