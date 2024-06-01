package domain

import "github.com/iarsham/teacher-tool-api/internal/models"

type PassUsecase interface {
	FindById(id uint64) (*models.Users, error)
	EncryptPass(plainPass string) ([]byte, error)
	UpdatePassword(id uint64, password string) (*models.Users, error)
}
