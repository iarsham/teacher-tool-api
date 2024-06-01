package domain

import (
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/models"
)

type UserRepository interface {
	FindAll() ([]*models.Users, error)
	FindById(id uint64) (*models.Users, error)
	FindByPhone(phone string) (*models.Users, error)
	Create(user *entities.UserRequest) (*models.Users, error)
	Update(id uint64, user *entities.UpdateUserRequest) (*models.Users, error)
	UpdatePassword(id uint64, password string) (*models.Users, error)
	Delete(id uint64) error
}

type UserUsecase interface {
	FindAll() ([]*models.Users, error)
	FindById(id uint64) (*models.Users, error)
	FindByPhone(phone string) (*models.Users, error)
	Update(id uint64, user *entities.UpdateUserRequest) (*models.Users, error)
	Delete(id uint64) error
}
