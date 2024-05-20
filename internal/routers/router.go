package routers

import (
	"database/sql"
	"github.com/iarsham/multiplexer"
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/iarsham/teacher-tool-api/internal/handlers"
	"go.uber.org/zap"
	"net/http"
)

const BasePathV1 string = "/api/v1"

func Routes(db *sql.DB, l *zap.Logger, cfg *configs.Config) http.Handler {
	mux := multiplexer.New(http.NewServeMux(), BasePathV1)
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.HttpMethodHandler)
	mux.HandleFunc("GET /healthcheck", handlers.HealthCheckHandler)
	return mux
}
