package usecase

import (
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"go.uber.org/zap"
)

type passUsecase struct {
	userRepository domain.UserRepository
	logger         *zap.Logger
}

func NewPassUsecase(userRepository domain.UserRepository, logger *zap.Logger) domain.PassUsecase {
	return &passUsecase{
		userRepository: userRepository,
		logger:         logger,
	}
}
func (p *passUsecase) FindById(id uint64) (*models.Users, error) {
	user, err := p.userRepository.FindById(id)
	if err != nil {
		p.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (p *passUsecase) EncryptPass(plainPass string) ([]byte, error) {
	encryptPass, err := helpers.EncryptPass(plainPass)
	if err != nil {
		p.logger.Error(err.Error())
		return []byte(nil), err
	}
	return encryptPass, nil
}

func (p *passUsecase) UpdatePassword(id uint64, password string) (*models.Users, error) {
	user, err := p.userRepository.UpdatePassword(id, password)
	if err != nil {
		p.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}
