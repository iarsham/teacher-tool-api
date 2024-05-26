package usecase

import (
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	userRepository domain.UserRepository
}

func NewLoginUsecase(userRepository domain.UserRepository) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
	}
}

func (a *loginUsecase) FindByPhone(phone string) (*models.Users, error) {
	return a.userRepository.FindByPhone(phone)
}

func (a *loginUsecase) VerifyPass(user *models.Users, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
