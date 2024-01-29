package routes

import (
	"orb-api/controllers"
	"github.com/gofiber/fiber/v2"
) 

func Setup(app *fiber.App, controllers *controllers.Controllers) {
	api := app.Group("/")

	api.Post("/register", controllers.User.CreateUser)
	api.Delete("/user/disable/:id", controllers.User.DeleteUser)
	api.Get("/sortprojects/:id", controllers.User.SortProjects)
}
