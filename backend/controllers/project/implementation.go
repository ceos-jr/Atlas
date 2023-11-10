package project

import (
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBody struct {
	Name: 		string `json:"name" validade:"required"`
	SectorID:	uint`json:"sectorid" validade:"required"`
	AdmID:		uint`json:"admid" validade:"required"`
}

func (handler *BaseHandler) CreateProject(context *fiber.Ctx) error{
	body := new(CreateUserRequestBody)

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

	newProject, serviceError := handler.Service.CreateProject(
		body.Name,
		body.SectorID,
		body.AdmID,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "project creation error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "project created successfully",
		"user":    newProject,
	})
}

