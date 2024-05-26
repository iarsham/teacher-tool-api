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

func Routes(db *sql.DB, logger *zap.Logger, cfg *configs.Config) http.Handler {
	mux := multiplexer.New(http.NewServeMux(), BasePathV1)
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.HttpMethodHandler)
	dynamic := multiplexer.NewChain(
		middlewares.LoggerMiddleware(logger),
		middlewares.RecoveryMiddleware(logger),
	)
	mux.Handle("GET /healthcheck", dynamic.WrapFunc(handlers.HealthCheckHandler))
	authGroup := mux.Group("/auth")
	AuthRouter(authGroup, dynamic, db, logger)
	return mux
}
