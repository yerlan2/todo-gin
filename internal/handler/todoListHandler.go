package handler

import (
	"net/http"
	"strconv"
	"yn/todo/internal/model"
	"yn/todo/internal/payload"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		// payload.NewErrorResponse(c, http.StatusUnauthorized,
		// 	err.Error(),
		// )
		return
	}
	lists, err := h.service.TodoList.GetAll(userId)
	if err != nil {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (h *Handler) getListById(c *gin.Context) {
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
	list, err := h.service.TodoList.GetById(userId, id)
	if err != nil {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		// payload.NewErrorResponse(c, http.StatusUnauthorized,
		// 	err.Error(),
		// )
		return
	}
	var input model.TodoList
	if err := c.BindJSON(&input); err != nil {
		payload.NewErrorResponse(c, http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	id, err := h.service.TodoList.Create(userId, input)
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

func (h *Handler) updateListById(c *gin.Context) {
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
	var input payload.UpdateTodoListRequest
	if err := c.BindJSON(&input); err != nil {
		payload.NewErrorResponse(c, http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	if err := h.service.TodoList.UpdateById(userId, id, input); err != nil {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	c.JSON(http.StatusOK, payload.StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) partUpdateListById(c *gin.Context) {
}

func (h *Handler) deleteListById(c *gin.Context) {
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
	err = h.service.TodoList.DeleteById(userId, id)
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
func (h *Handler) getListItemsById(c *gin.Context) {
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
	items, err := h.service.TodoList.GetAllItemsById(userId, id)
	if err != nil {
		payload.NewErrorResponse(c, http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	c.JSON(http.StatusOK, items)
}
