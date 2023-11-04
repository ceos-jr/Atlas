package user

import (
	"github.com/gofiber/fiber/v2"
	"orb-api/config"
	"orb-api/services/user"
)

type AuthHandler struct {
	AuthService *user.Service
	BaseHandler *user.BaseHandler
}

func NewAuthHandler(authService *user.Service, baseHandler *user.BaseHandler) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		BaseHandler: baseHandler,
	}
}

type LoginRequestBody struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (handler *AuthHandler) Login(context *fiber.Ctx) error {
	body := new(LoginRequestBody)

	if parseError := context.BodyParser(body); parseError != nil {
		return context.Status(fiber.tatusBadRequest).JSON(fiber.Map{
			"message":"Invalid request body",
			"errors": parseError.Error(),
		})
}
}

validationErrors := handler.BaseHandler.Validator.Validate(body)

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

token, authError := handler.AuthService.GenerateToken(body.Email, body.Password)

if authError != nil {
	return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Authentication failed",
		"error":   authError.Error(),
	})
}

return context.Status(fiber.StatusOK).JSON(fiber.Map{
	"message": "Authentication successful",
	"token":   token,
})

