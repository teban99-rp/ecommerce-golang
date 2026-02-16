package services

import (
	"errors"

	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/dto"
	"github.com/teban99-rp/ecommerce-golang/models"
	"github.com/teban99-rp/ecommerce-golang/utils"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUsers() ([]dto.UserResponseDTO, error)
	Login(email, password string) (string, error)
	FindByEmail(email string) (*models.User, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) CreateUser(user *models.User) error {

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashPassword

	return database.DB.Create(user).Error
}

func (s *userService) GetUsers() ([]dto.UserResponseDTO, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	var response []dto.UserResponseDTO
	for _, user := range users {
		response = append(response, dto.UserResponseDTO{
			ID:       user.ID,
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
		})
	}
	return response, nil
}

func (s *userService) Login(email, password string) (string, error) {

	user, err := s.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", errors.New("credenciales incorrectas")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
