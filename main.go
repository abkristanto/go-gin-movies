package main

import (
	"net/http"

	"github.com/abkristanto/go-gin-movies/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	h := handlers.InitHandler()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/movies", h.CreateMovie)
	r.POST("/register", h.RegisterUser)
	r.POST("/genres", h.CreateGenre)
	r.GET("/genres/:id", h.GetGenreById)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
