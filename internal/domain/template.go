package domain

import (
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"mime/multipart"
	"net/http"
)

type TemplateRepository interface {
	FindAll() ([]*models.Templates, error)
	Create(template *entities.TemplateRequest) (*models.Templates, error)
	Delete(templateID uint64) error
}

type TemplateUsecase interface {
	GetObjID(r *http.Request) (uint64, error)
	FindAll() ([]*models.Templates, error)
	Create(template *entities.TemplateRequest, userID uint64) (*models.Templates, error)
	Delete(templateID uint64) error
	UploadFile(file multipart.File, folder, fileName string) (string, error)
}
