package main

import (
	"github.com/iarsham/teacher-tool-api/internal/routers"
	"github.com/iarsham/teacher-tool-api/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	l := logger.NewZapLog()
	defer l.Sync()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      routers.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	l.Info("Starting server", zap.String("address", ":8080"))
	if err := srv.ListenAndServe(); err != nil {
		l.Fatal(err.Error())
	}
}
