package domain

import (
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"github.com/iarsham/teacher-tool-api/internal/models"
)

type RefreshTokenUsecase interface {
	ValidateRefreshToken(refreshToken string) (helpers.M, error)
	FindById(id uint64) (*models.Users, error)
	CreateAccessToken(userID uint64, phone string, role models.Role) (string, error)
}
