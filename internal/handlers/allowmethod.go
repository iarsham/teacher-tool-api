package handlers

import (
	"github.com/iarsham/bindme"
	"net/http"
)

func HttpMethodHandler(w http.ResponseWriter, r *http.Request) {
	err := `{"error": "The method specified in the request is not allowed for the resource. 
			Please check the allowed methods and try again."}`
	bindme.WriteJson(w, http.StatusMethodNotAllowed, err, nil)
}
