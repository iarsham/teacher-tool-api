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
	data.UserID = t.Usecase.GetUserID(r)
	file, handler, err := bindme.ReadFile(r, "file", maxFileSize)
	if err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	if _, err := t.Usecase.FindByFile(handler); !errors.Is(err, sql.ErrNoRows) {
		bindme.WriteJson(w, http.StatusConflict, helpers.M{"error": "template already exists"}, nil)
		return
	}
	helpers.Background(func() {
		link, _ := t.Usecase.UploadFile(file, "templates", handler.Filename)
		t.Usecase.Create(data, link)
	})
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
