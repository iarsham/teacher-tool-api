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
	mux.HandleFunc("GET /healthcheck", handlers.HealthCheckHandler)
	dynamic := multiplexer.NewChain(
		middlewares.LoggerMiddleware(logger),
		middlewares.RecoveryMiddleware(logger),
	)
	protected := dynamic.Append(
		middlewares.JwtAuthMiddleware(logger, cfg),
	)
	authGroup := mux.Group("/auth")
	userGroup := mux.Group("/user")
	templateGroup := mux.Group("/template")
	AuthRouter(authGroup, dynamic, db, logger, cfg)
	UserRouter(userGroup, protected, db, logger)
	TemplateRouter(templateGroup, protected, db, logger, cfg)
	return middlewares.CorsMiddleware(cfg).Handler(mux)
}
