package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (mr *AdRouter) AdBannerHandler(w http.ResponseWriter, r *http.Request) {
	padIDStr := r.URL.Query().Get("id")
	padID, err := strconv.Atoi(padIDStr)
	if err != nil {
		http.Error(w, "Invalid ad ID", http.StatusBadRequest)
		return
	}

	ad, err := mr.Ad.AdReadTarget(adID)
	fmt.Println(ad)
	uniqueLink := mr.addr + "/api/v1/redirect?id=" + adIDStr

	id := mr.SelectUC.Get()
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
