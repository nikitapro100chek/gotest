package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nikitapro100chek/gotest/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// api := router.Group("/api", h.userIdentity)
	// {
	lists := router.Group("/list", h.userIdentity)
	{
		lists.POST("/", h.createList)
		lists.GET("/", h.getAllLists)
	}
	edit := router.Group("/edit", h.userIdentity)
	{
		edit.POST("/:id", h.updateList)
	}

	return router
}
