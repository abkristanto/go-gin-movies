package services

import (
	"github.com/abkristanto/go-gin-movies/dtos"
	"github.com/abkristanto/go-gin-movies/models"
	g "github.com/abkristanto/go-gin-movies/repositories"
)

type GenreService interface {
	Create(genreRequest *dtos.GenreRequest) (*dtos.GenreResponse, error)
	GetGenre(id int) (*dtos.GenreResponse, error)
}

type genreService struct {
	genreRepository g.GenreRepository
}

type GSConfig struct {
	GenreRepository g.GenreRepository
}

func NewGenreService(c *GSConfig) GenreService {
	return &genreService{
		genreRepository: c.GenreRepository,
	}
}

func (g *genreService) Create(genreRequest *dtos.GenreRequest) (*dtos.GenreResponse, error) {
	genre := &models.Genre{
		Name: genreRequest.Name,
	}
	_, err := g.genreRepository.CreateGenre(genre)
	if err != nil {
		return nil, err
	}

	genreResponse := &dtos.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}
	return genreResponse, nil
}

func (g *genreService) GetGenre(id int) (*dtos.GenreResponse, error) {
	genre, err := g.genreRepository.QueryGenreById(id)
	if err != nil {
		return nil, err
	}
	genreResponse := &dtos.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}
	return genreResponse, nil
}
