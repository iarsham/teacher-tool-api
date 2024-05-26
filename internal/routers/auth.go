package routers

import (
	"database/sql"
	"github.com/iarsham/multiplexer"
	"github.com/iarsham/teacher-tool-api/internal/handlers"
	"github.com/iarsham/teacher-tool-api/internal/repository"
	"github.com/iarsham/teacher-tool-api/internal/usecase"
	"go.uber.org/zap"
)

func AuthRouter(r *multiplexer.Router, chain multiplexer.Chain, db *sql.DB, logger *zap.Logger) {
	userRepo := repository.NewUserRepository(db)
	hr := &handlers.RegisterHandler{
		Usecase: usecase.NewRegisterUsecase(userRepo),
		Logger:  logger,
	}
	r.Handle("POST /register", chain.WrapFunc(hr.RegisterHandler))
}
