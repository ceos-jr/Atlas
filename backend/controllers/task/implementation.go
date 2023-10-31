package task

import (
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBodyID struct {
	ID     uint `json:"id" validate:"required"`
}

//function that marks a task as concluded
func (handler *BaseHandler) ConcludeTask(context *fiber.Ctx) error {
	body := new(CreateUserRequestBodyID)

	//parsing the body
	if parseError := context.BodyParser(body); parseError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   parseError.Error(),
		})
	}

	//validating the body
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

	//excecuting the function ConcludedTask
	newTask, ServiceError := handler.Service.ConcludedTask(
		body.ID,
	)

	if ServiceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Task updating error",
			"error":   ServiceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "task marked as concluded",
		"user":    newTask,
	})
}