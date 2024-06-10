package handlers

import (
	"database/sql"
	"errors"
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"net/http"
)

type TemplateHandler struct {
	Usecase domain.TemplateUsecase
}

func (t *TemplateHandler) GetAllTemplatesHandler(w http.ResponseWriter, r *http.Request) {
	templates, err := t.Usecase.FindAll()
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusOK, templates, nil)
}

func (t *TemplateHandler) CreateTemplateHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entities.TemplateRequest)
	userID := helpers.GetUserID(r)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	file, handler, err := r.FormFile("file")
	data.File = handler
	if err != nil {
		switch {
		case errors.Is(err, http.ErrMissingFile):
			bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": "missing file in request"}, nil)
			return
		default:
			bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
			return
		}
	}
	defer file.Close()
	helpers.Background(func() {
		t.Usecase.UploadFile(file, "templates", handler.Filename)
	})
	if _, err := t.Usecase.Create(data, userID); err != nil {
		switch {
		case !errors.Is(err, sql.ErrNoRows):
			bindme.WriteJson(w, http.StatusConflict, helpers.M{"error": "template already exists"}, nil)
			return
		default:
			bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
			return
		}
	}
	bindme.WriteJson(w, http.StatusCreated, helpers.M{"response": "template created"}, nil)
}

func (t *TemplateHandler) DeleteTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateID, err := t.Usecase.GetObjID(r)
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	if err := t.Usecase.Delete(templateID); err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusNoContent, nil, nil)
}
