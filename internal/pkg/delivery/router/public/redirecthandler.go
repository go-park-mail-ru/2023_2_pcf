package router

import (
	"AdHub/internal/pkg/entities"
	"log"
	"net/http"
	"strconv"
)

func (mr *PublicRouter) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	Token := r.URL.Query().Get("id")
	if Token == "" {
		http.Error(w, "Ad token is missing", http.StatusBadRequest)
		return
	}

	padID := r.URL.Query().Get("pad")
	if padID == "" {
		http.Error(w, "Pad is missing", http.StatusBadRequest)
		return
	}

	adID, err := mr.ULink.GetAdId(Token)
	log.Printf("%s", adID)

	ad, err := mr.Ad.AdRead(adID)
	if err != nil {
		http.Error(w, "Ad is missing", http.StatusBadRequest)
		return
	}
	log.Printf("%s", ad.Id)
	mr.ULink.ULinkRemove(&entities.ULink{
		Token: Token,
		AdId:  adID,
	})

	pad_id, err := strconv.Atoi(padID)
	if err != nil {
		http.Error(w, "Padid cost to int error", http.StatusBadRequest)
		return
	}

	pad, err := mr.Pad.PadRead(pad_id)
	if err != nil {
		http.Error(w, "Pad read missing", http.StatusBadRequest)
		return
	}

	pad.Clicks += 1
	pad.Balance += ad.Click_cost
	err = mr.Pad.PadUpdate(pad)
	if err != nil {
		http.Error(w, "Pad update missing", http.StatusBadRequest)
		return
	}

	ad.Budget -= ad.Click_cost
	err = mr.Ad.AdUpdate(ad)
	if err != nil {
		http.Error(w, "Ad update missing", http.StatusBadRequest)
		return
	}

	website := "http://" + ad.Website_link

	http.Redirect(w, r, website, http.StatusSeeOther)
}
