package routers

import (
	"database/sql"
	"github.com/iarsham/multiplexer"
	"github.com/iarsham/teacher-tool-api/internal/handlers"
	"github.com/iarsham/teacher-tool-api/internal/repository"
	"github.com/iarsham/teacher-tool-api/internal/usecase"
	"go.uber.org/zap"
)

func UserRouter(r *multiplexer.Router, chain multiplexer.Chain, db *sql.DB, logger *zap.Logger) {
	userRepo := repository.NewUserRepository(db)
	hu := &handlers.UserHandler{
		Usecase: usecase.NewUserUsecase(userRepo, logger),
	}
	hp := &handlers.PasswordHandler{
		Usecase: usecase.NewPassUsecase(userRepo, logger),
	}
	r.Handle("GET ", chain.WrapFunc(hu.GetUserHandler))
	r.Handle("PUT ", chain.WrapFunc(hu.UpdateUserHandler))
	r.Handle("DELETE ", chain.WrapFunc(hu.DeleteUserHandler))
	r.Handle("POST /change-password", chain.WrapFunc(hp.PasswordChangeHandler))
}
