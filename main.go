package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/pkg/repository"
	service2 "github.com/heroku/go-getting-started/pkg/service"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		//log.Fatal("$PORT must be set")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://gvishnajhvjtem:52d1c94554ebcf6fb9979ffbd527688bf2ce061e67841b435a63d50997dd6884@ec2-52-45-211-119.compute-1.amazonaws.com:5432/d5dr1po1mfb1ds"
	}
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		repos := repository.NewRepository(db)
		service := service2.NewService(repos)
		id, err := service.Answer.CreateAnswer("тесирование 1")
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":    id,
			"error": err,
			"db":    db,
		})
	})

	router.POST("/auth/check", func(c *gin.Context) {
		var input signInInput

		if err := c.BindJSON(&input); err != nil {

		}

		repos := repository.NewRepository(db)
		service := service2.NewService(repos)

		id, _ := service.CheckAuth(input.phone)

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	})

	router.Run(":" + port)
}

type signInInput struct {
	phone string `json:"phone" binding:"required"`
}
