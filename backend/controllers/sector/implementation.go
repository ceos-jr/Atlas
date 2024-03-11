package sector

import "github.com/gofiber/fiber/v2"

type CreateSectorRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AdmID       uint   `json:"adm_id"`
}

func (handler *BaseHandler) CreateSector(context *fiber.Ctx) error {
	body := new(CreateSectorRequestBody)

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

	newUser, serviceError := handler.Service.CreateSector(
		body.Name,
		body.Description,
		body.AdmID,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "sector creation error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "sector created successfully",
		"user":    newUser,
	})
}
