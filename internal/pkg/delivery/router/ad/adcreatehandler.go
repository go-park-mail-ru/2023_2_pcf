package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/auth"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (mr *AdRouter) AdCreateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token       string `json:"token"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Sector      string `json:"sector"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	userId, err := auth.MySessionStorage.GetUserId(request.Token)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "Error get session", http.StatusBadRequest)
		return
	}

	ad := entities.Ad{
		Id:          1,
		Name:        request.Name,
		Description: request.Description,
		Sector:      request.Sector,
		Owner_id:    userId,
	}

	newAd, err := mr.Ad.AdCreate(&ad)
	if err != nil {
		http.Error(w, "Error create ad", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newAd)
	fmt.Println(len(auth.MySessionStorage.Sessions))
	fmt.Println(responseJSON)
	w.Write(responseJSON)
}
