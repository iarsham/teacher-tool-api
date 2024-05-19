package routers

import (
	"github.com/iarsham/multiplexer"
	"net/http"
)

const BasePathV1 string = "/api/v1"

func Routes() http.Handler {
	mux := multiplexer.New(http.NewServeMux(), BasePathV1)
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})
	return mux
}
