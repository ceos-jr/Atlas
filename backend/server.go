package main

import (
	"log"
	"orb-api/config"
	"orb-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	repository, setupError := config.SetupDB()
	
  if setupError != nil {
    log.Fatal(setupError)
  }

  handler := controllers.NewBaseHandler(repository)

	server.Get("/", handler.HandleHello)

	server.Listen(":8000")
}
