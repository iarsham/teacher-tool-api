package usecase

import (
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type registerUsecase struct {
	userRepository domain.UserRepository
	logger         *zap.Logger
}

func NewRegisterUsecase(userRepository domain.UserRepository, logger *zap.Logger) domain.RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (r *registerUsecase) FindByPhone(phone string) (*models.Users, error) {
	user, err := r.userRepository.FindByPhone(phone)
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (r *registerUsecase) Create(user *entities.UserRequest) (*models.Users, error) {
	u, err := r.userRepository.Create(user)
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	return u, nil
}

func (r *registerUsecase) EncryptPass(plainPass string) ([]byte, error) {
	encryptPass, err := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
	if err != nil {
		r.logger.Error(err.Error())
		return []byte(nil), err
	}
	return encryptPass, nil
}
