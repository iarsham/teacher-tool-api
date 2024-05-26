package main

import (
	"flag"
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/iarsham/teacher-tool-api/internal/database"
	"github.com/iarsham/teacher-tool-api/internal/routers"
	"github.com/iarsham/teacher-tool-api/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	logs := logger.NewZapLog(*debug)
	defer logs.Sync()

	cfg, err := configs.NewConfig()
	if err != nil {
		logs.Fatal(err.Error())
	}
	cfg.App.Debug = *debug

	db, err := database.OpenDB(cfg)
	if err != nil {
		logs.Fatal(err.Error())
	}
	defer db.Close()
	logs.Info("Database connected", zap.String("host", cfg.Postgres.Host), zap.Int("port", cfg.Postgres.Port))

	srv := &http.Server{
		Addr:         cfg.App.Addr,
		Handler:      routers.Routes(db, logs, cfg),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logs.Info("Starting server", zap.String("host", cfg.App.Host), zap.Int("port", cfg.App.Port))
	if err := srv.ListenAndServe(); err != nil {
		logs.Fatal(err.Error())
	}
}
