package usecase

import (
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	logger         *zap.Logger
	cfg            *configs.Config
}

func NewLoginUsecase(userRepository domain.UserRepository, logger *zap.Logger, cfg *configs.Config) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		logger:         logger,
		cfg:            cfg,
	}
}

func (l *loginUsecase) FindByPhone(phone string) (*models.Users, error) {
	user, err := l.userRepository.FindByPhone(phone)
	if err != nil {
		l.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (l *loginUsecase) VerifyPass(hashPass, plainPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(plainPass)); err != nil {
		l.logger.Error(err.Error())
		return err
	}
	return nil
}

func (l *loginUsecase) CreateAccessToken(userID uint64, phone string) (string, error) {
	access, err := helpers.CreateAccessToken(userID, phone, l.cfg.App.Secret, l.cfg.App.AccessHourTTL)
	if err != nil {
		l.logger.Error(err.Error())
		return "", err
	}
	return access, nil
}

func (l *loginUsecase) CreateRefreshToken(userID uint64) (string, error) {
	refresh, err := helpers.CreateRefreshToken(userID, l.cfg.App.Secret, l.cfg.App.RefreshHourTTL)
	if err != nil {
		l.logger.Error(err.Error())
		return "", err
	}
	return refresh, nil
}

func (l *loginUsecase) IsPhoneValid(phone string) bool {
	return helpers.IsPhoneValid(phone)
}
