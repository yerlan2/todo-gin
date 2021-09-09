package handler

import (
	"net/http"
	"yn/todo/internal/model"
	"yn/todo/internal/payload"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		payload.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		payload.NewErrorResponse(
			c, http.StatusInternalServerError, err.Error(),
		)
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input payload.SignInRequest
	if err := c.BindJSON(&input); err != nil {
		payload.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.Authorization.GenerateToken(
		*input.Username,
		*input.Password,
	)
	if err != nil {
		payload.NewErrorResponse(
			c, http.StatusInternalServerError, err.Error(),
		)
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}
