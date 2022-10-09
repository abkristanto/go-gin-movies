package handlers

import (
	"github.com/abkristanto/go-gin-movies/database"
	"github.com/abkristanto/go-gin-movies/repositories"
	"github.com/abkristanto/go-gin-movies/services"
)

func InitHandler() *Handler {
	database.Connect()
	ar := repositories.NewAuthRepository(&repositories.ARConfig{
		DB: database.Get(),
	})
	as := services.NewAuthService(&services.ASConfig{
		AuthRepository: ar,
	})
	gr := repositories.NewGenreRepository(&repositories.GRConfig{
		DB: database.Get(),
	})
	gs := services.NewGenreService(&services.GSConfig{
		GenreRepository: gr,
	})
	mr := repositories.NewMovieRepository(&repositories.MRConfig{
		DB: database.Get(),
	})
	ms := services.NewMovieService(&services.MSConfig{
		MovieRepository: mr,
		GenreRepository: gr,
	})
	h := New(&HandlerConfig{
		MovieService: ms,
		AuthService:  as,
		GenreService: gs,
	})
	return h
}
