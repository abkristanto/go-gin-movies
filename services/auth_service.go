package services

import (
	"github.com/abkristanto/go-gin-movies/dtos"
	"github.com/abkristanto/go-gin-movies/models"
	a "github.com/abkristanto/go-gin-movies/repositories"
)

type AuthService interface {
	Register(registerRequest *dtos.RegisterRequest) (*dtos.RegisterResponse, error)
}

type authService struct {
	authRepository a.AuthRepository
}

type ASConfig struct {
	AuthRepository a.AuthRepository
}

func NewAuthService(c *ASConfig) AuthService {
	return &authService{
		authRepository: c.AuthRepository,
	}
}

func (a *authService) Register(registerRequest *dtos.RegisterRequest) (*dtos.RegisterResponse, error) {
	user := &models.User{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}

	_, err := a.authRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	registerResponse := &dtos.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return registerResponse, nil
}
