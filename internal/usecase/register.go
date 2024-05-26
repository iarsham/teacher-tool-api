package usecase

import (
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type registerUsecase struct {
	userRepository domain.UserRepository
}

func NewRegisterUsecase(userRepository domain.UserRepository) domain.RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
	}
}

func (a *registerUsecase) FindByPhone(phone string) (*models.Users, error) {
	return a.userRepository.FindByPhone(phone)
}

func (a *registerUsecase) Create(user *entities.UserRequest) (*models.Users, error) {
	return a.userRepository.Create(user)
}

func (a *registerUsecase) EncryptPass(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 14)
}
