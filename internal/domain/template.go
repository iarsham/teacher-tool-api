package domain

import (
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"mime/multipart"
	"net/http"
)

type TemplateRepository interface {
	FindAll() ([]*models.Templates, error)
	Create(template *entities.TemplateRequest, link string) (*models.Templates, error)
	FindByFile(link string) (*models.Templates, error)
	Delete(templateID uint64) error
}

type TemplateUsecase interface {
	GetObjID(r *http.Request) (uint64, error)
	GetUserID(r *http.Request) uint64
	FindAll() ([]*models.Templates, error)
	FindByFile(file *multipart.FileHeader) (*models.Templates, error)
	Create(template *entities.TemplateRequest, link string) (*models.Templates, error)
	Delete(templateID uint64) error
	UploadFile(file multipart.File, folder, fileName string) (string, error)
}
