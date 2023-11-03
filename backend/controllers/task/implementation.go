package task

import (
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBodyId struct {
	Id     uint `json:"id" validate:"required"`
	AssignedTo  uint `json:"assigned_to" validate:"required"`
}

func (handler *BaseHandler) AssignedTo(context *fiber.Ctx) error {
	body := new(CreateUserRequestBodyId)

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

	newTask, serviceError := handler.Service.AssignTask(
		body.Id,
		body.AssignedTo,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Task updating error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task assigned successfuly",
		"task":    newTask,
	})
}
