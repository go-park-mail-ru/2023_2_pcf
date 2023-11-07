package router

import (
	"fmt"
	"net/http"
	"strconv"
)

func (mr *PublicRouter) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	adIDStr := r.URL.Query().Get("id")
	if adIDStr == "" {
		http.Error(w, "Ad ID is missing", http.StatusBadRequest)
		return
	}

	adID, err := strconv.Atoi(adIDStr)

	ad, err := mr.Ad.AdRead(adID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting ad URL: %v", err), http.StatusInternalServerError)
		return
	}
	website := "http://" + ad.Website_link

	http.Redirect(w, r, website, http.StatusSeeOther)
}
