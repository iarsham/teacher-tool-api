package middlewares

import (
	"github.com/iarsham/teacher-tool-api/pkg/response"
	"go.uber.org/zap"
	"net/http"
)

func RecoveryMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.Header().Set("Connection", "close")
					response.ServerErrJSON(w, logger, err.(error))
					return
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
