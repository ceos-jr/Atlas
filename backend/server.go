package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"orb-api/config"
	"orb-api/controllers"
)

func main() {
	server := fiber.New()

	repository, setupError := config.SetupDB(".env")

	if setupError != nil {
		log.Fatal(setupError)
	}

	defer config.CloseDB(repository)

	handler := controllers.NewBaseHandler(repository)

	server.Get("/", handler.HandleHello)

	server.Listen(":8000")
}
