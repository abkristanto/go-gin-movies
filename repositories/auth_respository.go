package repositories

import (
	"github.com/abkristanto/go-gin-movies/errors"
	"github.com/abkristanto/go-gin-movies/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) (*models.User, error)
}

type authRepository struct {
	database *gorm.DB
}

type ARConfig struct {
	DB *gorm.DB
}

func NewAuthRepository(c *ARConfig) AuthRepository {
	return &authRepository{
		database: c.DB,
	}
}

func (a *authRepository) CreateUser(user *models.User) (*models.User, error) {
	result := a.database.Where(&models.User{Email: user.Email}).Or(&models.User{Username: user.Username}).FirstOrCreate(user)
	if result.RowsAffected == 0 {
		return nil, &errors.DuplicateUserError{}
	}
	return user, nil
}
