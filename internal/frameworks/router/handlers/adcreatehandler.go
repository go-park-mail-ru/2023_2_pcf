package router

import (
	"AdHub/internal/app/frameworks/auth"
	"AdHub/internal/app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *APIServer) AdCreateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8081")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8081")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

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

	ad := models.Ad{
		Id:          1,
		Name:        request.Name,
		Description: request.Description,
		Sector:      request.Sector,
		Owner_id:    userId,
	}

	newUser, err := s.Store.Ad().Create(&ad)
	if err != nil {
		http.Error(w, "Error create ad", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Status - 201
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(newUser)
	fmt.Println(len(auth.MySessionStorage.Sessions))
	fmt.Println(responseJSON)
	w.Write(responseJSON)
}
