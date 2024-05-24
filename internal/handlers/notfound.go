package handlers

import (
	"github.com/iarsham/bindme"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := `{"error": "The requested resource could not be found. Please check the URL and try again."}`
	bindme.WriteJson(w, http.StatusNotFound, err, nil)
}
