package usecase

import (
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"go.uber.org/zap"
)

type refreshTokenUsecaseImpl struct {
	userRepository domain.UserRepository
	logger         *zap.Logger
	cfg            *configs.Config
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, logger *zap.Logger, cfg *configs.Config) domain.RefreshTokenUsecase {
	return &refreshTokenUsecaseImpl{
		userRepository: userRepository,
		logger:         logger,
		cfg:            cfg,
	}
}

func (u *refreshTokenUsecaseImpl) ValidateRefreshToken(refreshToken string) (helpers.M, error) {
	isValid, err := helpers.IsTokenValid(refreshToken, u.cfg.App.Secret)
	if err != nil {
		u.logger.Error("failed to validate refresh token", zap.Error(err))
		return nil, err
	}
	claims, err := helpers.GetClaims(isValid)
	if err != nil {
		u.logger.Error("failed to get claims from refresh token", zap.Error(err))
		return nil, err
	}
	return claims, nil
}

func (u *refreshTokenUsecaseImpl) FindById(id uint64) (*models.Users, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		u.logger.Error("failed to find user by id", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (u *refreshTokenUsecaseImpl) CreateAccessToken(userID uint64, phone string, role models.Role) (string, error) {
	accessToken, err := helpers.CreateAccessToken(userID, phone, role, u.cfg.App.Secret, u.cfg.App.AccessHourTTL)
	if err != nil {
		u.logger.Error("failed to create access token", zap.Error(err))
		return "", err
	}
	return accessToken, nil
}
