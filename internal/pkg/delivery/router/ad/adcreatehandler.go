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
	request.Click_cost = r.FormValue("click_cost")

	uidAny := r.Context().Value("userId")
	uid, ok := uidAny.(int)
	if !ok {
		mr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
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

	newclickcost, err := strconv.ParseFloat(request.Click_cost, 64)
	if err != nil {
		mr.logger.Error("Error click cost parse" + err.Error())
		http.Error(w, "Error click cost parse", http.StatusInternalServerError)
		return
	}

	user, err := mr.User.UserReadById(uid)
	if err != nil {
		mr.logger.Error("Error retrieving user: " + err.Error())
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	currentBalance, err := mr.Balance.BalanceRead(user.BalanceId)
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
		Name:         request.Name,
		Description:  request.Description,
		Website_link: request.WebsiteLink,
		Budget:       newbudget,
		Image_link:   image,
		Owner_id:     uid,
		Target_id:    target,
		Click_cost:   newclickcost,
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
