package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/proto/api"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func (mr *PublicRouter) GetBanner(w http.ResponseWriter, r *http.Request) {
	padIDStr := r.URL.Query().Get("id")
	padID, err := strconv.Atoi(padIDStr)
	if err != nil {
		http.Error(w, "Invalid pad ID", http.StatusBadRequest)
		return
	}

	pad, err := mr.Pad.PadRead(padID)
	if err != nil {
		http.Error(w, "Invalid pad parse", http.StatusBadRequest)
		return
	}

	padtarget, err := mr.Target.TargetRead(pad.Target_id)

	targets, err := mr.Target.TargetRandom()
	if err != nil {
		http.Error(w, "Invalid targets get", http.StatusBadRequest)
		return
	}

	selectRequests := &api.SelectRequests{}

	// Пройдите по массиву targets и создайте экземпляры SelectRequest для каждого элемента
	for _, target := range targets {
		selectRequest := &api.SelectRequest{
			Id:        int64(target.Id),
			Name:      target.Name,
			OwnerId:   int64(target.Owner_id),
			Gender:    target.Gender,
			MinAge:    int64(target.Min_age),
			MaxAge:    int64(target.Max_age),
			Interests: target.Interests,
			Tags:      target.Tags,
			Keys:      target.Keys,
			Regions:   target.Regions,
		}

		// Добавьте созданный экземпляр SelectRequest в массив requests
		selectRequests.Requests = append(selectRequests.Requests, selectRequest)
	}

	selectRequests.Id = int64(padtarget.Id)
	selectRequests.Name = padtarget.Name
	selectRequests.OwnerId = int64(padtarget.Owner_id)
	selectRequests.Gender = padtarget.Gender
	selectRequests.MinAge = int64(padtarget.Min_age)
	selectRequests.MaxAge = int64(padtarget.Max_age)
	selectRequests.Interests = padtarget.Interests
	selectRequests.Tags = padtarget.Tags
	selectRequests.Keys = padtarget.Keys
	selectRequests.Regions = padtarget.Regions

	id, err := mr.SelectUC.Get(context.Background(), selectRequests)
	if err != nil {
		http.Error(w, "Invalid get ad", http.StatusBadRequest)
		return
	}

	ads, err := mr.Ad.AdByTarget(int(id.Id))
	if err != nil {
		http.Error(w, "Invalid get ad by target", http.StatusBadRequest)
		return
	}

	ad := ads[0]
	token := uuid.New().String()

	mr.ULink.ULinkCreate(&entities.ULink{
		Token: token,
		AdId:  0,
	})

	uniqueLink := mr.addr + "/api/v1/redirect?id=" + token + "?pad=" + strconv.Itoa(pad.Id)
	data := struct {
		Link     string
		ImageURL string
	}{
		Link:     "http://" + uniqueLink,
		ImageURL: mr.addr + "/api/v1/file?file=" + ad.Image_link,
	}

	tmpl := "<a href=\"" + data.Link + "\"><img src=\"" + data.ImageURL + "\" alt=\"Ad Banner\"></a>"

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(tmpl)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
