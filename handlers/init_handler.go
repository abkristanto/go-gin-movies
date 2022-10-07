package handlers

import (
	"github.com/abkristanto/go-gin-movies/database"
	"github.com/abkristanto/go-gin-movies/repositories"
	"github.com/abkristanto/go-gin-movies/services"
)

func InitHandler() *Handler {
	database.Connect()
	mr := repositories.NewMovieRepository(&repositories.MRConfig{
		DB: database.Get(),
	})
	ms := services.NewMovieService(&services.MSConfig{
		MovieRepository: mr,
	})
	ar := repositories.NewAuthRepository(&repositories.ARConfig{
		DB: database.Get(),
	})
	as := services.NewAuthService(&services.ASConfig{
		AuthRepository: ar,
	})
	h := New(&HandlerConfig{
		MovieService: ms,
		AuthService:  as,
	})
	return h
}
