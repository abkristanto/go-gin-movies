package handlers

import (
	"net/http"

	"github.com/abkristanto/go-gin-movies/dtos"
	"github.com/abkristanto/go-gin-movies/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMovie(c *gin.Context) {
	var createMovieRequest *dtos.CreateMovieRequest

	err := c.ShouldBindJSON(&createMovieRequest)
	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	
	movieResponse, err := h.movieService.CreateMovie(createMovieRequest)

	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteSuccessResponse(c, movieResponse)
}
