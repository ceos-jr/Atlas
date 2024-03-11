package user_role

import (
	"github.com/gofiber/fiber/v2"
)

type CreateUserRoleRequestBodyId struct {
	Id      uint `json:"id"`
	UserId	uint `json:"user_id" validate:"required"`
	RoleId	uint `json:"role_id" validate:"required"`
}

func (handler *BaseHandler) AssignedRole(context *fiber.Ctx) error {
	body := new(CreateUserRoleRequestBodyId)

	if parseError := context.BodyParser(body); parseError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   parseError.Error(),
		})
	}

	validationErrors := handler.Validator.Validate(body)

	if validationErrors != nil {
		errorMessages := make([]string, len(validationErrors))

		for index := range validationErrors {
			errorMessages[index] = validationErrors[index].Message
		}

		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"errors":  errorMessages,
		})
	}

	UserRole, serviceError := handler.Service.AssigneRole(
		body.UserId,
		body.RoleId,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Role assigned error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Role assigned successfully",
		"user_role":    UserRole,
	})
}

func (handler *BaseHandler) UnassignRole(context *fiber.Ctx) error {
	body := new(CreateUserRoleRequestBodyId)

	if parseError := context.BodyParser(body); parseError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   parseError.Error(),
		})
	}

	validationErrors := handler.Validator.Validate(body)

	if validationErrors != nil {
		errorMessages := make([]string, len(validationErrors))

		for index := range validationErrors {
			errorMessages[index] = validationErrors[index].Message
		}

		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"errors":  errorMessages,
		})
	}

	UserRole, serviceError := handler.Service.UnassignRole(
		body.UserId,
		body.RoleId,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Role unassign error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Role unassign successfully",
		"user_role":    UserRole,
	})
}