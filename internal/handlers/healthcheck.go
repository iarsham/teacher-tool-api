package handlers

import (
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"net/http"
)

// HealthCheckHandler godoc
//
//	@Summary	Health check
//	@Tags		Server
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.HealthCheck
//	@Router		/healthcheck [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := helpers.M{
		"status": "available",
	}
	bindme.WriteJson(w, http.StatusOK, helpers.M{"response": data}, nil)
}
