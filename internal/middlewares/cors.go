package middlewares

import (
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/rs/cors"
	"net/http"
)

func CorsMiddleware(cfg *configs.Config) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: cfg.App.CorsOrigins,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            cfg.App.Debug,
		MaxAge:           cfg.App.CorsMaxAge,
	})
}
