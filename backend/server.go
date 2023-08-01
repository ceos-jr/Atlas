package main

import (
  "github.com/gofiber/fiber/v2"
  "orb-api/controllers"
  "orb-api/config"
)

func main() {
  server := fiber.New()

  repository := config.SetupDB() 
  handler := controllers.NewBaseHandler(repository)
  
  server.Get("/", handler.HandleHello)

  server.Listen(":8000")
}
