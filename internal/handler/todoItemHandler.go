package handler

import (
	"net/http"
	"strconv"
	"yn/todo/internal/payload"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		// payload.NewErrorResponse(c, http.StatusUnauthorized,
		// 	err.Error(),
		// )
		return
	}
	items, err := h.service.TodoItem.GetAll(userId)
	if err != nil {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		// payload.NewErrorResponse(c, http.StatusUnauthorized,
		// 	err.Error(),
		// )
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		payload.NewErrorResponse(c, http.StatusBadRequest,
			"invalid id param",
		)
		return
	}
	item, err := h.service.TodoItem.GetById(userId, id)
	if err != nil {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		// payload.NewErrorResponse(c, http.StatusUnauthorized,
		// 	err.Error(),
		// )
		return
	}
	var input payload.TodoItemRequest
	if err := c.BindJSON(&input); err != nil {
		payload.NewErrorResponse(c, http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	id, err := h.service.TodoItem.Create(userId, input)
	if err != nil {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		// payload.NewErrorResponse(c, http.StatusUnauthorized,
		// 	err.Error(),
		// )
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		payload.NewErrorResponse(c, http.StatusBadRequest,
			"invalid id param",
		)
		return
	}
	var input payload.UpdateTodoItemRequest
	if err := c.BindJSON(&input); err != nil {
		payload.NewErrorResponse(c, http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	if err := h.service.TodoItem.UpdateById(userId, id, input); err != nil {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	c.JSON(http.StatusOK, payload.StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) partUpdateItemById(c *gin.Context) {
}

func (h *Handler) deleteItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		// payload.NewErrorResponse(c, http.StatusUnauthorized,
		// 	err.Error(),
		// )
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		payload.NewErrorResponse(c, http.StatusBadRequest,
			"invalid id param",
		)
		return
	}
	err = h.service.TodoItem.DeleteById(userId, id)
	if err != nil {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	c.JSON(http.StatusOK, payload.StatusResponse{
		Status: "ok",
	})
}
