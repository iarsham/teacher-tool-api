package middlewares

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
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
				bindme.WriteJson(w, http.StatusUnauthorized, helpers.M{"error": "Authorization header not provided"}, nil)
				return
			}
			authToken := strings.Split(authHeader, " ")
			if len(authToken) != 2 || strings.ToLower(authToken[0]) != "bearer" {
				bindme.WriteJson(w, http.StatusUnauthorized, helpers.M{"response": "Invalid Authorization header"}, nil)
				return
			}
			token, err := helpers.IsTokenValid(authToken[1], cfg.App.Secret)
			if err != nil {
				switch {
				case errors.Is(err, jwt.ErrTokenExpired):
					bindme.WriteJson(w, http.StatusUnauthorized, helpers.M{"error": "token is expired"}, nil)
				default:
					logger.Error("invalid token: failed to parse token", zap.Any("error", err))
					bindme.WriteJson(w, http.StatusUnauthorized, helpers.M{"error": "token is invalid"}, nil)
				}
				return
			}
			claims, err := helpers.GetClaims(token)
			if err != nil {
				logger.Error("invalid claims: failed to parse token", zap.Any("error", err))
				bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
				return
			}
			if claims["sub"] != nil {
				bindme.WriteJson(w, http.StatusUnprocessableEntity, helpers.M{"error": "refresh token not allowed"}, nil)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), "user_id", claims["user_id"]))
			r = r.WithContext(context.WithValue(r.Context(), "phone", claims["phone"]))
			r = r.WithContext(context.WithValue(r.Context(), "role", claims["role"]))
			next.ServeHTTP(w, r)
		})
	}
}

func IsAdminMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Context().Value("role") != "admin" {
				bindme.WriteJson(w, http.StatusForbidden, helpers.M{"error": "forbidden, just admin allowed"}, nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
