package usecase

import (
	"context"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepository domain.UserRepository
	ctx            context.Context
}

func NewAuthUsecase(userRepository domain.UserRepository) domain.AuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
	}
}

func (a *authUsecase) FindByPhone(phone string) (*models.Users, error) {
	return a.userRepository.FindByPhone(phone)
}

func (a *authUsecase) Create(user *entities.UserRequest) (*models.Users, error) {
	return a.userRepository.Create(user)
}

func (a *authUsecase) EncryptPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (a *authUsecase) VerifyPass(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}
