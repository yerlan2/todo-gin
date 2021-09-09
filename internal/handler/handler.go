package handler

import (
	"yn/todo/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.GET("/:id/items", h.getListItemsById)
			lists.POST("/", h.createList)
			lists.PUT("/:id", h.updateListById)
			lists.PATCH("/:id", h.partUpdateListById)
			lists.DELETE("/:id", h.deleteListById)
		}
		items := api.Group("/items")
		{
			items.GET("/", h.getAllItems)
			items.GET("/:id", h.getItemById)
			items.POST("/", h.createItem)
			items.PUT("/:id", h.updateItemById)
			items.PATCH("/:id", h.partUpdateItemById)
			items.DELETE("/:id", h.deleteItemById)
		}
	}

	return router
}
