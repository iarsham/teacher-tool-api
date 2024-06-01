package routers

import (
	"database/sql"
	"github.com/iarsham/multiplexer"
	_ "github.com/iarsham/teacher-tool-api/api"
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/iarsham/teacher-tool-api/internal/handlers"
	"github.com/iarsham/teacher-tool-api/internal/middlewares"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
)

func Routes(db *sql.DB, logger *zap.Logger, cfg *configs.Config) http.Handler {
	mux := multiplexer.New(http.NewServeMux(), cfg.App.BaseAPI)
	mux.Handle("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.HttpMethodHandler)
	dynamic := multiplexer.NewChain(
		middlewares.LoggerMiddleware(logger),
		middlewares.RecoveryMiddleware(logger),
	)
	mux.Handle("GET /healthcheck", dynamic.WrapFunc(handlers.HealthCheckHandler))
	authGroup := mux.Group("/auth")
	AuthRouter(authGroup, dynamic, db, logger, cfg)
	protected := dynamic.Append(
		middlewares.JwtAuthMiddleware(logger, cfg),
	)
	userGroup := mux.Group("/user")
	UserRouter(userGroup, protected, db, logger, cfg)
	return middlewares.CorsMiddleware(cfg).Handler(mux)
}
