package handlers

import "net/http"

func HttpMethodHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(`{"error": "The method specified in the request is not allowed for the resource. Please check the allowed methods and try again."}`))
}
