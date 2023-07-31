package controllers

import (
	"log"
	"orb-api/config"

	"github.com/gofiber/fiber/v2"
)

type BaseHandler struct {
  Repository *config.Repository
}

func NewBaseHandler(repository *config.Repository) *BaseHandler {
  return &BaseHandler{
    Repository: repository,
  }
}

func (handler *BaseHandler) HandleHello(context *fiber.Ctx) error {
  if handler.Repository.DB == nil {
    log.Fatal("Missing Database connection")
  }

  return context.SendString("Orb API is running")
} 
