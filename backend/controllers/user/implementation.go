package user

import (
	"orb-api/repositories/user"

	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBody struct {
	ID       uint   `json:"ID" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Status   uint   `json:"status" validate:"required"`
}



func (handler *BaseHandler) CreateUser(context *fiber.Ctx) error {
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

	newUser, serviceError := handler.Service.CreateUser(
		body.Name,
		body.Email,
		"12345678",
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user creation error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
		"user":    newUser,
	})
}

// func (handler *BaseHandler) ReadUser(context *fiber.Ctx) error {
// 	body := new(CreateUserRequestBody)

// 	if parseError := context.BodyParser(body); parseError != nil {
// 		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid request body",
// 			"error":   parseError.Error(),
// 		})
// 	}

// 	validationErrors := handler.Validator.Validate(body)

// 	if validationErrors != nil {
// 		errorMessages := make([]string, len(validationErrors))

// 		for index := range validationErrors {
// 			errorMessages[index] = validationErrors[index].Message
// 		}

// 		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid request body",
// 			"errors":  errorMessages,
// 		})
// 	}

// 	ReadUser, serviceError := handler.Service.ReadUser(user.IReadBy{
// 		ID: &body.ID,
// 		Name: &body.Name,
// 		Email: &body.Email,
// 		Status: &body.Status,
// 	})

// 	if serviceError != nil {
// 		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "user read error",
// 			"error":   serviceError.Error(),
// 		})
// 	}

// 	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"message": "user read successfully",
// 		"user":    ReadUser,
// 	})
// }

// func (handler *BaseHandler) UpdateUser(context *fiber.Ctx) error {
// 	body := new(CreateUserRequestBody)

// 	if parseError := context.BodyParser(body); parseError != nil {
// 		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid request body",
// 			"error":   parseError.Error(),
// 		})
// 	}

// 	validationErrors := handler.Validator.Validate(body)

// 	if validationErrors != nil {
// 		errorMessages := make([]string, len(validationErrors))

// 		for index := range validationErrors {
// 			errorMessages[index] = validationErrors[index].Message
// 		}

// 		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid request body",
// 			"errors":  errorMessages,
// 		})
// 	}

// 	UpdateUser, serviceError := handler.Service.UpdateUser(user.IUpdate{
// 		ID: body.ID,
// 		Name: &body.Name,
// 		Email: &body.Email,
// 		Password: &body.Password,
// 		Status: &body.Status,
	
// 	})
		

// 	if serviceError != nil {
// 		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "user update error",
// 			"error":   serviceError.Error(),
// 		})
// 	}

// 	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"message": "user update successfully",
// 		"user":    UpdateUser,
// 	})
// }

func (handler *BaseHandler) DeleteUser(context *fiber.Ctx) error {
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

	DeleteUser, serviceError := handler.Service.DeleteUser(
		body.ID,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user delete error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user deleted successfully",
		"user":    DeleteUser,
	})
}




