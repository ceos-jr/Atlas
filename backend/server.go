package main

import (
	"log"
	"orb-api/config"
	"orb-api/controllers"
	"orb-api/services"

	"github.com/gofiber/fiber/v2"
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

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("oiii")
	})
	server.Post("/register", controllers.User.CreateUser)

	/* Usuario: */

	// Disable user:
	server.Delete("/user/disable/:id", controllers.User.DeleteUser)

	server.Listen(":8000")
}
