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
  if repository == nil {
    log.Fatal("Missing Database connection")
  }

  return &BaseHandler{
    Repository: repository,
  }
}

func (handler *BaseHandler) HandleHello(context *fiber.Ctx) error {
  return context.SendString("Orb API is running")
} 
