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
	r.GET("/movies", h.GetMovies)
	r.POST("/register", h.RegisterUser)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
