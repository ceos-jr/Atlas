package user

import (
	"github.com/gofiber/fiber/v2"
	userService "orb-api/services/user"
	"strconv"
)

type Controller struct {
	UserService userService.Interface
}

func New(userService userService.Interface) *Controller {
	return &Controller{UserService: userService}
}

func (uc *Controller) Create(c *fiber.Ctx) error {
	var createUser userService.ICreateUser
	if err := c.BodyParser(&createUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newUser, err := uc.UserService.CreateNewUser(createUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func (uc *Controller) UpdateName(c *fiber.Ctx) error {
	userID := c.Params("id")
	var update userService.IUpdate

	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	updatedUser, err := uc.UserService.UpdateName(uint(id), *update.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

func (uc *Controller) UpdateEmail(c *fiber.Ctx) error {
	panic("Not implemented yet")
}

func (uc *Controller) UpdatePassword(c *fiber.Ctx) error {
	panic("Not implemented yet")
}

func (uc *Controller) UpdateStatus(c *fiber.Ctx) error {
	panic("Not implemented yet")
}
