package router

import (
	"bytes"
	"net/http"
	"strings"
	"time"
)

func (mr *UserRouter) FileHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	fileData, err := mr.File.Get(filename)
	if err != nil {
		mr.logger.Error("Error getting file: " + err.Error())
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}

	if strings.HasSuffix(filename, ".png") {
		w.Header().Set("Content-Type", "image/png")
	} else if strings.HasSuffix(filename, ".jpg") {
		w.Header().Set("Content-Type", "image/jpeg")
	}

	http.ServeContent(w, r, filename, time.Now(), bytes.NewReader(fileData))
}
