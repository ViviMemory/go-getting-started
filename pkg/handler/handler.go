package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/check", h.authCheck)
	}

	return router
}
