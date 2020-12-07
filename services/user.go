package services

import (
	"time"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/0xkalvin/url-shortener/models"
	"github.com/0xkalvin/url-shortener/repositories"
	"gopkg.in/lucsky/cuid.v1"
)

//UserService  interface
type UserService struct {
	Repository repositories.UserRepository
}

// NewUserService creates an service layer
func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{
		Repository: repository,
	}
}

// CreateUser creates an user
func (service *UserService) CreateUser(name string, email string) (*models.User, error) {
	logger := log.GetLogger()

	user, err := service.Repository.Create(&models.User{
		ID:        cuid.New(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})

	if err != nil {
		logger.Error("Failed to create user entity")

		return nil, err
	}

	return user, nil
}

// FindOneUser returns an user if it exists
func (service *UserService) FindOneUser(id string) (*models.User, error) {
	logger := log.GetLogger()

	user, err := service.Repository.FindOne(id)

	if err != nil {
		logger.Error("Failed to find user entity")

		return nil, err
	}

	return user, nil
}
