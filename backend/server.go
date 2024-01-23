package main

import (
	"log"
	"orb-api/config"
	"orb-api/controllers"
	"orb-api/routes"
	"orb-api/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	server := fiber.New()

	server.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000",
        AllowHeaders: "Origin, Content-Type, Accept",
        AllowCredentials: true,
    }))

	repository, setupError := config.SetupDB(".env")

	if setupError != nil {
		log.Fatal(setupError)
	}
	defer config.CloseDB(repository)

	services := services.SetupServices(repository)

	controllers := controllers.SetupControllers(services)

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("oii")
	})

	routes.Setup(server, controllers)
	server.Listen(":8000")
}