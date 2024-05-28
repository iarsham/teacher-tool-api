package handlers

import (
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"net/http"
)

func HttpMethodHandler(w http.ResponseWriter, r *http.Request) {
	data := "The method specified in the request is not allowed for the resource." +
		"Please check the allowed methods and try again."
	bindme.WriteJson(w, http.StatusMethodNotAllowed, helpers.M{"error": data}, nil)
}
