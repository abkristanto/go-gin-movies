package repositories

import "gorm.io/gorm"

type MovieRepository interface {
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
