package domain

import (
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"mime/multipart"
	"net/http"
)

type QuestionsRepository interface {
	FindAll() ([]*models.Questions, error)
	FindByFile(link string) (*models.Questions, error)
	FindByID(id uint64) (*models.Questions, error)
	Create(question *entities.QuestionRequest, link string) (*models.Questions, error)
	Delete(id uint64) error
}

type QuestionsUsecase interface {
	GetObjID(r *http.Request) (uint64, error)
	GetUserID(r *http.Request) uint64
	FindAll() ([]*models.Questions, error)
	FindByFile(file *multipart.FileHeader) (*models.Questions, error)
	FindByID(id uint64) (*models.Questions, error)
	Create(question *entities.QuestionRequest, link string) (*models.Questions, error)
	Delete(id uint64) error
	UploadFile(file multipart.File, folder, fileName string) (string, error)
}
