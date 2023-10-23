package user

import (
	"orb-api/repositories/user"

	"github.com/gofiber/fiber/v2"
)

type UpdateUserRequestBody struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Status uint  `json:"status" validate:"required"`
}

func (handler *BaseHandler) UpdateUserName(context *fiber.Ctx) error {
	body := new(UpdateUserRequestBody)

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

	newUpdate, serviceError := handler.Service.UpdateUser(user.IUpdate{
		ID:       body.ID,
		Name:     &body.Name,
		Email:    &body.Email,
		Password: &body.Password,
		Status:   &body.Status,
	})

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
