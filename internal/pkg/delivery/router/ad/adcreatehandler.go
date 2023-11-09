package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func (mr *AdRouter) AdCreateHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		//Token       string  `json:"token"`
		Name        string `json:"name"`
		Description string `json:"description"`
		WebsiteLink string `json:"website_link"`
		Budget      string `json:"budget"`
		TargetId    string `json:"target_id"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		mr.logger.Error("Invalid request body" + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	uidAny := r.Context().Value("userId")
	uid, ok := uidAny.(int)
	if !ok {
		mr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

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

	//userId, err := mr.Session.GetUserId(request.Token)
	//if err != nil {
	//	mr.logger.Error("Error get session" + err.Error())
	//	http.Error(w, "Error get session", http.StatusBadRequest)
	//	return
	//}
	newbudget, err := strconv.ParseFloat(request.Budget, 64)
	if err != nil {
		mr.logger.Error("Error budget parse" + err.Error())
		http.Error(w, "Error budget parse", http.StatusInternalServerError)
		return
	}

	currentBalance, err := mr.Balance.BalanceRead(uid)
	if currentBalance.Available_balance < newbudget {
		mr.logger.Error("Недостаточно средства" + err.Error())
		http.Error(w, "Недостаточно средств", http.StatusMethodNotAllowed)
		return
	} else {
		err = mr.Balance.BalanceReserve(newbudget, currentBalance.Id)
		if err != nil {
			mr.logger.Error("Error balance update" + err.Error())
			http.Error(w, "Error balance update", http.StatusInternalServerError)
			return
		}
	}

	target, err := strconv.Atoi(request.TargetId)
	if err != nil {
		mr.logger.Error("Error target parse" + err.Error())
		http.Error(w, "Error target parse", http.StatusInternalServerError)
		return
	}

	ad := entities.Ad{
		//Id:           1,
		Name:         request.Name,
		Description:  request.Description,
		Website_link: request.WebsiteLink,
		Budget:       newbudget, // Преобразование int в float64
		Image_link:   image,     // Используйте Imagelink из request
		Owner_id:     uid,       // Укажите нужное значение Owner_id
		Target_id:    target,    // Укажите нужное значение Target_id
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
