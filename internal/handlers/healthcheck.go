package handlers

import (
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := helpers.M{
		"status": "available",
	}
	helpers.WriteJson(w, http.StatusOK, data)
}
