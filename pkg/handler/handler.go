package handler

import "github.com/heroku/go-getting-started/pkg/service"

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
