package controllers

import (
	"orb-api/repositories"
	"github.com/gofiber/fiber/v2"
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
