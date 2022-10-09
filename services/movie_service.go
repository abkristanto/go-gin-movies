package services

import (
	"github.com/abkristanto/go-gin-movies/dtos"
	"github.com/abkristanto/go-gin-movies/models"
	m "github.com/abkristanto/go-gin-movies/repositories"
)

type MovieService interface {
	CreateMovie(createMovieRequest *dtos.CreateMovieRequest) (*dtos.CreateMovieResponse, error)
	DeleteMovie(id int) error
}

type movieService struct {
	movieRepository m.MovieRepository
	genreRepository m.GenreRepository
}

type MSConfig struct {
	MovieRepository m.MovieRepository
	GenreRepository m.GenreRepository
}

func NewMovieService(c *MSConfig) MovieService {
	return &movieService{
		movieRepository: c.MovieRepository,
		genreRepository: c.GenreRepository,
	}
}

func (m *movieService) CreateMovie(createMovieRequest *dtos.CreateMovieRequest) (*dtos.CreateMovieResponse, error) {
	var genreResponses []*dtos.GenreResponse
	var genres []*models.Genre
	for _, genreID := range createMovieRequest.Genres {
		genre, err := m.genreRepository.QueryGenreById(genreID)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
		genreResponse := &dtos.GenreResponse{
			ID:   genre.ID,
			Name: genre.Name,
		}
		genreResponses = append(genreResponses, genreResponse)
	}
	movie := &models.Movie{
		Title:    createMovieRequest.Title,
		Image:    createMovieRequest.Image,
		Synopsis: createMovieRequest.Synopsis,
		Genres:   genres,
	}

	_, err := m.movieRepository.CreateMovie(movie)
	if err != nil {
		return nil, err
	}

	movieResponse := &dtos.CreateMovieResponse{
		ID:       movie.ID,
		Title:    movie.Title,
		Image:    movie.Image,
		Synopsis: movie.Synopsis,
		Genres:   genreResponses,
	}

	return movieResponse, nil
}

func (m *movieService) DeleteMovie(id int) error {
	err := m.movieRepository.DeleteMovie(id)
	if err != nil {
		return err 
	}
	return nil
}
