package main

import (
  "github.com/gofiber/fiber/v2"
)

func main() {
  server := fiber.New()

  server.Get("/", func(context *fiber.Ctx) error {
    return context.SendString("Orb API is running")
  })

  server.Listen(":8000")
}
