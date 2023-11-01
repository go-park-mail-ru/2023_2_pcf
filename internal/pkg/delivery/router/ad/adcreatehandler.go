package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"net/http"
)

func (mr *AdRouter) AdCreateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token       string  `json:"token"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		WebsiteLink string  `json:"website_link"`
		Budget      float64 `json:"budget"`
		TargetId    int     `json:"target_id"`
		ImageLink   string  `json:"image_Link"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		mr.logger.Error("Invalid request body" + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	userId, err := mr.Session.GetUserId(request.Token)
	if err != nil {
		mr.logger.Error("Error get session" + err.Error())
		http.Error(w, "Error get session", http.StatusBadRequest)
		return
	}

	ad := entities.Ad{
		//Id:           1,
		Name:         request.Name,
		Description:  request.Description,
		Website_link: request.WebsiteLink,
		Budget:       request.Budget,    // Преобразование int в float64
		Image_link:   request.ImageLink, // Используйте Imagelink из request
		Owner_id:     userId,            // Укажите нужное значение Owner_id
		Target_id:    request.TargetId,  // Укажите нужное значение Target_id
	}

	newAd, err := mr.Ad.AdCreate(&ad)
	if err != nil {
		mr.logger.Error("Error ad create" + err.Error())
		http.Error(w, "Error create ad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newAd)
	w.Write(responseJSON)
}
