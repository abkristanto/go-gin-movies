package handlers

import (
	"net/http"
	"strconv"

	"github.com/abkristanto/go-gin-movies/dtos"
	"github.com/abkristanto/go-gin-movies/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGenre(c *gin.Context) {
	var genreRequest *dtos.GenreRequest

	err := c.ShouldBindJSON(&genreRequest)
	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	genreResponse, err := h.genreService.Create(genreRequest)

	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	helpers.WriteSuccessResponse(c, genreResponse)
}

func (h *Handler) GetGenreById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	genreResponse, err := h.genreService.GetGenre(id)
	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	helpers.WriteSuccessResponse(c, genreResponse)

}
