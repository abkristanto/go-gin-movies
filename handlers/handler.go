package handlers

import "github.com/abkristanto/go-gin-movies/services"

type Handler struct {
	movieService services.MovieService
	authService  services.AuthService
}

type HandlerConfig struct {
	MovieService services.MovieService
	AuthService  services.AuthService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		movieService: c.MovieService,
		authService:  c.AuthService,
	}
}
