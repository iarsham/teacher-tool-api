package handlers

import (
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"github.com/iarsham/teacher-tool-api/pkg/response"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := helpers.M{
		"status": "available",
	}
	response.JSON(w, http.StatusOK, nil, data, nil)
}
