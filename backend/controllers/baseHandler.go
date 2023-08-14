package controllers

import (
	"github.com/gofiber/fiber/v2"
	"orb-api/repositories"
)

type BaseHandler struct {
	Repository *repository.Repository
}

func NewBaseHandler(repository *repository.Repository) *BaseHandler {
	return &BaseHandler{
		Repository: repository,
	}
}

func (handler *BaseHandler) HandleHello(context *fiber.Ctx) error {
	return context.SendString("Orb API is running")
}
