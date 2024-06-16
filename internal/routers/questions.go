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

func questionsRouter(r *multiplexer.Router, chain multiplexer.Chain, db *sql.DB, logger *zap.Logger, cfg *configs.Config) {
	questionRepo := repository.NewQuestionRepository(db)
	hq := &handlers.QuestionHandler{
		Usecase: usecase.NewQuestionsUsecase(questionRepo, logger, cfg),
	}
	r.HandleFunc("GET ", hq.GetAllQuestionsHandler)
	r.HandleFunc("GET /{id}", hq.GetQuestionHandler)
	r.Handle("POST ", chain.WrapFunc(hq.CreateQuestionHandler))
	r.Handle("DELETE /{id}", chain.WrapFunc(hq.DeleteQuestionHandler))
}
