package repositories

import (
	"fmt"
	"github.com/abkristanto/go-gin-movies/models"
	"gorm.io/gorm"
)

type GenreRepository interface {
	QueryGenreById(id int) (*models.Genre, error)
	CreateGenre(genre *models.Genre) (*models.Genre, error)
}

type genreRepository struct {
	database *gorm.DB
}

type GRConfig struct {
	DB *gorm.DB
}

func NewGenreRepository(c *GRConfig) GenreRepository {
	return &genreRepository{
		database: c.DB,
	}
}

func (g *genreRepository) CreateGenre(genre *models.Genre) (*models.Genre, error) {
	result := g.database.Create(&genre)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return genre, nil
}

func (g *genreRepository) QueryGenreById(id int) (*models.Genre, error) {
	fmt.Println("id")
	fmt.Println("finding...")
	var genre *models.Genre
	result := g.database.Find(&genre, id)
	fmt.Println((result.RowsAffected))
	fmt.Println(genre)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return genre, nil
}
