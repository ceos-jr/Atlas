package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"orb-api/config"
	"orb-api/controllers"
	"orb-api/services"
)

func main() {
	server := fiber.New()

	repository, setupError := config.SetupDB(".env")

	if setupError != nil {
		log.Fatal(setupError)
	}

	defer config.CloseDB(repository)

	services := services.SetupServices(repository)

	controllers := controllers.SetupControllers(services)

	server.Post("/register", controllers.User.CreateUser)

	server.Listen(":8000")
}
