package routes

import (
	"github.com/gofiber/fiber/v2"
	"orb-api/controllers"
)

func Setup(app *fiber.App, controllers *controllers.Controllers) {
	api := app.Group("/")

	api.Post("/register", controllers.User.CreateUser)
	api.Delete("/user/disable/:id", controllers.User.DeleteUser)
	api.Get("/sortprojects/:id", controllers.User.SortProjects)
	api.Put("/user/update/:id", controllers.User.UpdateUser)
	api.Get("/listusers", controllers.User.ReadAllUsers)
	api.Get("/user/:id", controllers.User.ReadUser)

	api.Post("/role/create", controllers.Role.CreateRole)
	api.Get("/role/listroles", controllers.Role.ReadAllRoles)
	api.Post("/user_role/atribuir", controllers.User_Role.AssignedRole)
	api.Post("/user_role/retirar", controllers.User_Role.UnassignRole)
}
