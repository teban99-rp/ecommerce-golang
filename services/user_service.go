package services

import (
	"github.com/teban99-rp/ecommerce-golang/models"
	"github.com/teban99-rp/ecommerce-golang/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUsers() ([]models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.repo.FindAll()
}
