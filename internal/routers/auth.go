package routers

import (
	"database/sql"
	"github.com/iarsham/multiplexer"
	"github.com/iarsham/teacher-tool-api/internal/handlers"
	"github.com/iarsham/teacher-tool-api/internal/repository"
	"github.com/iarsham/teacher-tool-api/internal/usecase"
	"go.uber.org/zap"
)

func AuthRouter(r *multiplexer.Router, chain multiplexer.Chain, db *sql.DB, l *zap.Logger) {
	userRepo := repository.NewUserRepository(db)
	h := &handlers.AuthHandler{
		AuthUsecase: usecase.NewAuthUsecase(userRepo),
	}
	r.Handle("POST ", chain.WrapFunc(h.LoginRegisterHandler))
}
