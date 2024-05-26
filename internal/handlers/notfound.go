package handlers

import (
	"github.com/iarsham/teacher-tool-api/pkg/response"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := "The requested resource could not be found. Please check the URL and try again."
	response.ErrJSON(w, http.StatusNotFound, nil, err)
}
