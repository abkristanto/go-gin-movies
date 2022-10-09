package main

import (
	"net/http"

	"github.com/abkristanto/go-gin-movies/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	h := handlers.InitHandler()
	api := r.Group("/api")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		genres := api.Group("/genres")
		{
			genres.POST("/", h.CreateGenre)
			genres.GET("/:id", h.GetGenreById)
		}
		api.POST("/movies", h.CreateMovie)
		api.DELETE("/movies/:id", h.DeleteMovie)
		api.POST("/register", h.RegisterUser)
	}
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
