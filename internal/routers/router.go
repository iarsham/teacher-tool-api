package routers

import (
	"github.com/iarsham/multiplexer"
	"github.com/iarsham/teacher-tool-api/internal/handlers"
	"net/http"
)

const BasePathV1 string = "/api/v1"

func Routes() http.Handler {
	mux := multiplexer.New(http.NewServeMux(), BasePathV1)
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.HttpMethodHandler)
	return mux
}
