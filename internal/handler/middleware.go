package handler

import (
	"errors"
	"net/http"
	"strings"
	"yn/todo/internal/payload"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		payload.NewErrorResponse(c, http.StatusUnauthorized,
			"empty auth header",
		)
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		payload.NewErrorResponse(c, http.StatusUnauthorized,
			"invalid auth header",
		)
		return
	}
	// parse token
	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		payload.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			"user id not found",
		)
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			"user id is of invalid type",
		)
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}
