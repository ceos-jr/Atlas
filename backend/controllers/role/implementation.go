package role

import (
	"github.com/gofiber/fiber/v2"
)
type CreateRoleRequestBody struct {
	ID       uint   `json:"ID"`
	Name     string `json:"name" validate:"required"`
	Description    string `json:"description" validate:"required"`
}

func (handler *BaseHandler) CreateRole(context *fiber.Ctx) error {
	body := new(CreateRoleRequestBody)

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

	newRole, serviceError := handler.Service.NewRole(
		body.Name,
		body.Description,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Role creation error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Role created successfully",
		"user":    newRole,
	})
}