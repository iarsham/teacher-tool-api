package domain

import "github.com/iarsham/teacher-tool-api/internal/models"

type LoginUsecase interface {
	FindByPhone(phone string) (*models.Users, error)
	VerifyPass(hashPass, plainPass string) error
	CreateAccessToken(userID uint64, phone string) (string, error)
	CreateRefreshToken(userID uint64) (string, error)
}
