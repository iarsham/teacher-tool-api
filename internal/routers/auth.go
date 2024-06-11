package routers

import (
	"database/sql"
	"github.com/iarsham/multiplexer"
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/iarsham/teacher-tool-api/internal/handlers"
	"github.com/iarsham/teacher-tool-api/internal/repository"
	"github.com/iarsham/teacher-tool-api/internal/usecase"
	"go.uber.org/zap"
)

func authRouter(r *multiplexer.Router, chain multiplexer.Chain, db *sql.DB, logger *zap.Logger, cfg *configs.Config) {
	userRepo := repository.NewUserRepository(db)
	hr := &handlers.RegisterHandler{
		Usecase: usecase.NewRegisterUsecase(userRepo, logger),
	}
	hl := &handlers.LoginHandler{
		Usecase: usecase.NewLoginUsecase(userRepo, logger, cfg),
	}
	r.Handle("POST /register", chain.WrapFunc(hr.RegisterHandler))
	r.Handle("POST /login", chain.WrapFunc(hl.LoginHandler))
}
