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
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
	}

	router.POST("/company", h.getCompany)

	group := router.Group("/group", h.userIdentity)
	{
		group.POST("/create", h.createGroup)
		group.POST("/adduser", h.addUser)
		//group.POST("/deleteuser", h.deleteUser)
		group.GET("/list", h.allGroup)
		//group.POST("/detail", h.detailGroup)

	}

	return router
}
