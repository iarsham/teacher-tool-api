package routers

import (
	"database/sql"
	"github.com/iarsham/multiplexer"
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/iarsham/teacher-tool-api/internal/handlers"
	"github.com/iarsham/teacher-tool-api/internal/middlewares"
	"go.uber.org/zap"
	"net/http"
)

const BasePathV1 string = "/api/v1"

func Routes(db *sql.DB, l *zap.Logger, cfg *configs.Config) http.Handler {
	mux := multiplexer.New(http.NewServeMux(), BasePathV1)
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.HttpMethodHandler)
	dynamic := multiplexer.NewChain(
		middlewares.LoggerMiddleware(l),
		middlewares.RecoveryMiddleware(l),
	)
	mux.Handle("GET /healthcheck", dynamic.WrapFunc(handlers.HealthCheckHandler))
	return mux
}
