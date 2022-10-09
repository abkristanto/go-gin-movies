package handlers

import (
	"net/http"
	"strconv"

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

func (h *Handler) DeleteMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	err = h.movieService.DeleteMovie(id)
	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.WriteSuccessResponse(c, nil)
}
