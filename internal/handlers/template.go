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

// GetAllTemplatesHandler godoc
//
//	@Summary		Get Templates
//	@Description	Get all templates
//	@Produce		json
//	@Tags			Templates
//	@Success		200	{object}	response.AllTemplates
//	@Failure		500	{object}	response.InternalServerError
//	@router			/template [get]
func (t *TemplateHandler) GetAllTemplatesHandler(w http.ResponseWriter, r *http.Request) {
	templates, err := t.Usecase.FindAll()
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusOK, templates, nil)
}

// CreateTemplateHandler godoc
//
//	@Summary		Create Template
//	@Description	Create a new template for exam
//	@Accept			multipart/form-data
//	@Produce		json
//	@Tags			Templates
//	@Param			templateRequest	body		entities.TemplateRequest	true	"Template data"
//	@Success		201				{object}	response.TemplateCreated
//	@Failure		400				{object}	response.BadRequest
//	@Failure		408				{object}	response.TemplateExists
//	@Failure		500				{object}	response.InternalServerError
//	@router			/template [post]
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

// DeleteTemplateHandler godoc
//
//	@Summary		Delete Template
//	@Description	Delete a template with id
//	@Produce		json
//	@Tags			Templates
//	@Param			id	path		int	true	"Template ID"
//	@Success		204	{object}	nil
//	@Failure		500	{object}	response.InternalServerError
//	@router			/template/{id} [delete]
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
