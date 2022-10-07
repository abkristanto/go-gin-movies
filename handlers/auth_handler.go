package handlers

import (
	"net/http"

	"github.com/abkristanto/go-gin-movies/dtos"
	"github.com/abkristanto/go-gin-movies/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var registerRequest *dtos.RegisterRequest

	err := c.ShouldBindJSON(&registerRequest)
	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	userResponse, err := h.authService.Register(registerRequest)

	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteSuccessResponse(c, userResponse)
}
