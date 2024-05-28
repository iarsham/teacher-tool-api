package handlers

import (
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := "The requested resource could not be found. Please check the URL and try again."
	bindme.WriteJson(w, http.StatusNotFound, helpers.M{"error": data}, nil)
}
