package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/pkg/service"
	memorycache "github.com/maxchagin/go-memorycache-example"
	"time"
)

type Handler struct {
	services *service.Service
	cache    *memorycache.Cache
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
		cache:    memorycache.New(36600*time.Hour, 36600*time.Hour),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		sms := auth.Group("/sms")
		{
			sms.POST("/send", h.sendSms)
			sms.POST("/check", h.checkSms)
		}

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
