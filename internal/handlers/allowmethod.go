package handlers

import (
	"github.com/iarsham/teacher-tool-api/pkg/response"
	"net/http"
)

func HttpMethodHandler(w http.ResponseWriter, r *http.Request) {
	err := "The method specified in the request is not allowed for the resource." +
		"Please check the allowed methods and try again."
	response.ErrJSON(w, http.StatusMethodNotAllowed, nil, err)
}
