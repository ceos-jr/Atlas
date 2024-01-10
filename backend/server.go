package main

import (
	"github.com/gofiber/fiber/v2"
	"orb-api/routes"
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

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Chablau!!")
	})

	routes.Setup(server, controllers)
	server.Listen(":8000")
}
