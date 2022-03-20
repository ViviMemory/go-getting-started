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
		auth.POST("/check", h.authCheck)
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
	}

	main := router.Group("/main", h.userIdentity)
	{
		main.GET("/info", h.info)
		main.GET("/setrole", h.setRoleAdmin)
		main.POST("/group/detail", h.detailGroup)
		main.POST("/group/invite", h.inviteUserInGroup)
		main.GET("/group/list", h.listInviteUser)
		main.POST("/group/active", h.activeInviteUser)
		main.GET("/test/categories/list", h.testCategoriesList)
		main.POST("/test/categories/created", h.testCategoriesAdd)
		main.POST("/test/created", h.CreatedTest)
		main.GET("/test/all", h.AllTests)
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
