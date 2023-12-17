package router

import (
	"io"
	"net/http"
	"strconv"
)

func (mr *AdRouter) AdUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		//Token       string   `json:"token"`
		AdId        string `json:"ad_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		WebsiteLink string `json:"website_link"`
		Budget      string `json:"budget"`
		TargetId    string `json:"target_id"`
		Click_cost  string `json:"click_cost"`
	}

	var image string
	err := r.ParseMultipartForm(100 << 20) // 10 MB - максимальный размер файла
	if err == nil {
		file, Header, err := r.FormFile("image")
		if err == nil {
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
	}

	err = r.ParseForm()
	if err != nil {
		mr.logger.Error("Error parsing form: " + err.Error())
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	request.Name = r.FormValue("name")
	request.Description = r.FormValue("description")
	request.WebsiteLink = r.FormValue("website_link")
	request.Budget = r.FormValue("budget")
	request.TargetId = r.FormValue("target_id")
	request.AdId = r.FormValue("ad_id")
	request.Click_cost = r.FormValue("click_cost")

	uidAny := r.Context().Value("userId")

	uid, ok := uidAny.(int)
	if !ok {
		mr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	newbudget, err := strconv.ParseFloat(request.Budget, 64)
	if err != nil {
		mr.logger.Error("Error budget parse" + err.Error())
		http.Error(w, "Error budget parse", http.StatusInternalServerError)
		return
	}

	newclickcost, err := strconv.ParseFloat(request.Click_cost, 64)
	if err != nil {
		mr.logger.Error("Error click cost parse" + err.Error())
		http.Error(w, "Error click cost parse", http.StatusInternalServerError)
		return
	}

	target, err := strconv.Atoi(request.TargetId)
	if err != nil {
		mr.logger.Error("Error target parse" + err.Error())
		http.Error(w, "Error target parse", http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(request.AdId)
	if err != nil {
		mr.logger.Error("Error ad parse" + err.Error())
		http.Error(w, "Error ad parse", http.StatusInternalServerError)
		return
	}

	// Получение текущего состояния рекламы из базы данных
	currentAd, err := mr.Ad.AdRead(id)
	if err != nil {
		mr.logger.Error("Error retrieving ad: " + err.Error())
		http.Error(w, "Error retrieving ad", http.StatusInternalServerError)
		return
	}

	// Проверяем, что реклама принадлежит пользователю из сессии
	if currentAd.Owner_id != uid {
		mr.logger.Error("User does not have permission to update this ad")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Обновляем поля рекламы
	if request.Name != "" {
		currentAd.Name = request.Name
	}
	if request.Description != "" {
		currentAd.Description = request.Description
	}
	if request.WebsiteLink != "" {
		currentAd.Website_link = request.WebsiteLink
	}
	if request.Budget != "" {
		currentAd.Budget = newbudget
	}
	if request.TargetId != "" {
		currentAd.Target_id = target
	}
	if image != "" {
		currentAd.Image_link = image
	}
	if request.Click_cost != "" {
		currentAd.Click_cost = newclickcost
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
