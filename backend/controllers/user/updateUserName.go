package user

import "github.com/gofiber/fiber/v2"

type UpdateNameRequestBody struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

func (handler *BaseHandler) UpdateUserName(context *fiber.Ctx) error {
	body := new(UpdateNameRequestBody)

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
			"message": "invalid request body",
			"errors":  errorMessages,
		})
	}

	newUpdate, serviceError := handler.Service.UpdateName(
		body.ID,
		body.Name,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"user update name error",
			"user": newUpdate,
		})
	}
	return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message":"user name update successfully",
		"user": newUpdate,
	})
}
