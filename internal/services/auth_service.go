package services

import (
	"errors"

	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/pkg/utils"
)

type AuthService struct {
	userService *UserService
	jwtSecret   string
}

func NewAuthService(userService *UserService, jwtSecret string) *AuthService {
	return &AuthService{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

// Login validates credentials and returns user + token
func (s *AuthService) Login(email, password string) (*models.User, string, error) {
	user, err := s.userService.GetByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Role.Name, s.jwtSecret)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) GetUserByID(id uint) (*models.User, error) {
	return s.userService.GetByID(id)
}
