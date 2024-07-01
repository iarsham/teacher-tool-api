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

// GetAllQuestionsHandler godoc
//
//	@Summary		Get All Questions
//	@Description	Get all questions handled by teacher
//	@Accept			json
//	@Produce		json
//	@Tags			Questions
//	@Param			page		query		int		false	"Page number"
//	@Param			page_size	query		int		false	"Page size"
//	@Success		200	{object}	response.AllQuestions
//	@Success		400	{object}	response.BadRequest
//	@Failure		500	{object}	response.InternalServerError
//	@router			/question [get]
func (q *QuestionHandler) GetAllQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	var filter helpers.PaginateFilter
	v := bindme.New()
	qs := r.URL.Query()
	filter.Page = v.ReadInt(qs, "page", 1)
	filter.PageSize = v.ReadInt(qs, "page_size", 10)
	if filter.Validate(v); !v.IsValid() {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": v.Errors}, nil)
		return
	}
	questions, metaData, err := q.Usecase.FindAll(filter.Limit(), filter.OffSet())
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusOK, helpers.M{"metadata": metaData, "questions": questions}, nil)
}

// GetQuestionHandler godoc
//
//	@Summary		Get Question
//	@Description	Get single question with id
//	@Accept			json
//	@Produce		json
//	@Tags			Questions
//	@Param			id	path		int	true	"Question ID"
//	@Success		200	{object}	response.QuestionData
//	@Failure		400	{object}	response.BadRequest
//	@Failure		404	{object}	response.QuestionNotFound
//	@Failure		500	{object}	response.InternalServerError
//	@router			/question/:id [get]
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

// CreateQuestionHandler godoc
//
//	@Summary		Create Question
//	@Description	Create a new question
//	@Accept			multipart/form-data
//	@Produce		json
//	@Tags			Questions
//	@Param			questionRequest	body		entities.QuestionRequest	true	"Question data"
//	@Success		200				{object}	response.QuestionCreated
//	@Failure		400				{object}	response.BadRequest
//	@Failure		408				{object}	response.QuestionAlreadyExists
//	@Failure		500				{object}	response.InternalServerError
//	@router			/question [post]
func (q *QuestionHandler) CreateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	question := new(entities.QuestionRequest)
	question.UserID = q.Usecase.GetUserID(r)
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
		q.Usecase.Create(question, link)
	})
	bindme.WriteJson(w, http.StatusCreated, helpers.M{"response": "question created"}, nil)
}

// DeleteQuestionHandler godoc
//
//	@Summary		Delete Question
//	@Description	Delete question with id
//	@Accept			json
//	@Produce		json
//	@Tags			Questions
//	@Param			id	path		int	true	"Question ID"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	response.BadRequest
//	@Failure		500	{object}	response.InternalServerError
//	@router			/question/:id [delete]
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
