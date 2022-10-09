package repositories

import (
	"github.com/abkristanto/go-gin-movies/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MovieRepository interface {
	CreateMovie(movie *models.Movie) (*models.Movie, error)
	DeleteMovie(id int) error
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

func (m *movieRepository) DeleteMovie(id int) error {
	var movie *models.Movie
	err := m.database.Find(&movie, id).Select(clause.Associations).Delete(&movie).Error
	if err != nil {
		return err
	}
	return nil
}
