package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
)

const maxFileSize = 2

type QuestionHandler struct {
	Usecase domain.QuestionsUsecase
}

func (q *QuestionHandler) GetAllQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	questions, err := q.Usecase.FindAll()
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusOK, questions, nil)
}

func (q *QuestionHandler) GetQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := q.Usecase.GetObjID(r)
	if err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	question, err := q.Usecase.FindByID(id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			bindme.WriteJson(w, http.StatusNotFound, helpers.M{"error": "question not found"}, nil)
		default:
			bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		}
		return
	}
	bindme.WriteJson(w, http.StatusOK, helpers.M{"response": question}, nil)
}

func (q *QuestionHandler) CreateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	userID := q.Usecase.GetUserID(r)
	question := new(entities.QuestionRequest)
	file, handler, err := bindme.ReadFile(r, "file", maxFileSize)
	if err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	if err := bindme.ReadForm(r, question); err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	if _, err := q.Usecase.FindByFile(handler); !errors.Is(err, sql.ErrNoRows) {
		bindme.WriteJson(w, http.StatusConflict, helpers.M{"error": "question already exists"}, nil)
		return
	}
	helpers.Background(func() {
		link, _ := q.Usecase.UploadFile(file, "questions", handler.Filename)
		q.Usecase.Create(question, link, userID)
	})
	bindme.WriteJson(w, http.StatusCreated, helpers.M{"response": "question created"}, nil)
}

func (q *QuestionHandler) DeleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := q.Usecase.GetObjID(r)
	if err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	if err := q.Usecase.Delete(id); err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusNoContent, nil, nil)
}
