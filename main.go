package main

import (
	"github.com/heroku/go-getting-started/pkg/handler"
	"github.com/heroku/go-getting-started/pkg/repository"
	service2 "github.com/heroku/go-getting-started/pkg/service"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://gvishnajhvjtem:52d1c94554ebcf6fb9979ffbd527688bf2ce061e67841b435a63d50997dd6884@ec2-52-45-211-119.compute-1.amazonaws.com:5432/d5dr1po1mfb1ds"
	}
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	repos := repository.NewRepository(db)
	service := service2.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(handler.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
