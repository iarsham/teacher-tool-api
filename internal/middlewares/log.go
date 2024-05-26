package middlewares

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
	return
}

func LoggerMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrapped := wrapResponseWriter(w)
			defer func() {
				start := time.Now()
				logger.Info(
					"Request processed",
					zap.Int("status", wrapped.status),
					zap.String("domain", r.Host),
					zap.String("method", r.Method),
					zap.String("path", r.URL.EscapedPath()),
					zap.Duration("time", time.Since(start)),
				)
			}()
			next.ServeHTTP(wrapped, r)
		})
	}
}
