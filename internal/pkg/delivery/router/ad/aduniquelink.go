package router

import (
	"net/http"
	"strconv"
	"text/template"
)

func (mr *AdRouter) AdBannerHandler(w http.ResponseWriter, r *http.Request) {
	adIDStr := r.URL.Query().Get("id")
	adID, err := strconv.Atoi(adIDStr)
	if err != nil {
		http.Error(w, "Invalid ad ID", http.StatusBadRequest)
		return
	}

	ad, err := mr.Ad.AdRead(adID)

	uniqueLink := mr.addr + "api/v1/redirect?id=" + adIDStr

	tmpl, err := template.New("ad_template").Parse(`<a href="{{.Link}}"><img src="{{.ImageURL}}" alt="Ad Banner"></a>`)
	if err != nil {
		http.Error(w, "Failed to create template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Link     string
		ImageURL string
	}{
		Link:     uniqueLink,
		ImageURL: mr.addr + "/api/v1/file?file=" + ad.Image_link,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)
}
