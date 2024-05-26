package domain

import "github.com/iarsham/teacher-tool-api/internal/models"

type LoginUsecase interface {
	FindByPhone(phone string) (*models.Users, error)
	VerifyPass(user *models.Users, password string) error
}
