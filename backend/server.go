package main

import (
	"github.com/gofiber/fiber/v2"
	"orb-api/config"
	"orb-api/controllers"
)

func main() {
	server := fiber.New()

	repository, _ := config.SetupDB()
	handler := controllers.NewBaseHandler(repository)

	server.Get("/", handler.HandleHello)

	server.Listen(":8000")
}
