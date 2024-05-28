package middlewares

import (
	"context"
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func JwtAuthMiddleware(logger *zap.Logger, cfg *configs.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				bindme.WriteJson(w, http.StatusUnauthorized, "Authorization header not provided", nil)
				return
			}
			authToken := strings.Split(authHeader, " ")
			if len(authToken) != 2 || strings.ToLower(authToken[0]) != "bearer" {
				bindme.WriteJson(w, http.StatusUnauthorized, "Invalid Authorization header", nil)
				return
			}
			claims, err := helpers.ExtractToken(authToken[1], cfg.App.Secret)
			if err != nil {
				logger.Error(err.Error())
				bindme.WriteJson(w, http.StatusInternalServerError, helpers.ErrInternalServer.Error(), nil)
				return
			}
			if claims["sub"] != nil {
				bindme.WriteJson(w, http.StatusUnprocessableEntity, "refresh token not allowed", nil)
			}
			r = r.WithContext(context.WithValue(r.Context(), "user", claims))
			next.ServeHTTP(w, r)
		})
	}
}
