package usecase

import (
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"go.uber.org/zap"
)

type userUsecase struct {
	userRepository domain.UserRepository
	logger         *zap.Logger
}

func NewUserUsecase(userRepository domain.UserRepository, logger *zap.Logger) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (u *userUsecase) FindAll() ([]*models.Users, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) FindById(id uint64) (*models.Users, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) FindByPhone(phone string) (*models.Users, error) {
	user, err := u.userRepository.FindByPhone(phone)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Update(id uint64, req *entities.UpdateUserRequest) (*models.Users, error) {
	updatedUser, err := u.userRepository.Update(id, req)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return updatedUser, nil
}

func (u *userUsecase) Delete(id uint64) error {
	err := u.userRepository.Delete(id)
	if err != nil {
		u.logger.Error(err.Error())
		return err
	}
	return nil
}
