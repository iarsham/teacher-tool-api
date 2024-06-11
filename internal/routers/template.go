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

func templateRouter(r *multiplexer.Router, chain multiplexer.Chain, db *sql.DB, logger *zap.Logger, cfg *configs.Config) {
	tmplRepo := repository.NewTemplateRepository(db)
	ht := &handlers.TemplateHandler{
		Usecase: usecase.NewTemplateUsecase(tmplRepo, logger, cfg),
	}
	r.Handle("GET ", chain.WrapFunc(ht.GetAllTemplatesHandler))
	r.Handle("POST ", chain.WrapFunc(ht.CreateTemplateHandler))
	r.Handle("DELETE /{id}", chain.WrapFunc(ht.DeleteTemplateHandler))
}
