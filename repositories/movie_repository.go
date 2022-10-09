package repositories

import (
	"github.com/abkristanto/go-gin-movies/models"
	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(movie *models.Movie) (*models.Movie, error)
}

type movieRepository struct {
	database *gorm.DB
}

type MRConfig struct {
	DB *gorm.DB
}

func NewMovieRepository(c *MRConfig) MovieRepository {
	return &movieRepository{
		database: c.DB,
	}
}

func (m *movieRepository) CreateMovie(movie *models.Movie) (*models.Movie, error) {
	result := m.database.Create(&movie)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return movie, nil
}
