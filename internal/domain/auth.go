package domain

import (
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
)

type AuthUsecase interface {
	FindByPhone(phone string) (*models.Users, error)
	Create(user *entities.UserRequest) (*models.Users, error)
	EncryptPass(password string) (string, error)
	VerifyPass(hash, plain string) bool
}
