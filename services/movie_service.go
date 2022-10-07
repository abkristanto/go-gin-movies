package services

import m "github.com/abkristanto/go-gin-movies/repositories"

type MovieService interface {
}

type movieService struct {
	movieRepository m.MovieRepository
}

type MSConfig struct {
	MovieRepository m.MovieRepository
}

func NewMovieService(c *MSConfig) MovieService {
	return &movieService{
		movieRepository: c.MovieRepository,
	}
}
