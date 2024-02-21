package user

import (
	"orb-api/repositories/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBody struct {
	ID       uint   `json:"ID" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Status   uint   `json:"status" validate:"required"`
}
type UpdateUserRequestBody struct {
	ID       uint    `json:"ID"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Status   *uint   `json:"status"`
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
		body.Password,
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

func (handler *BaseHandler) ReadUser(context *fiber.Ctx) error {
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

	ReadUser, serviceError := handler.Service.ReadUser(user.IReadBy{
		ID:     &body.ID,
		Name:   &body.Name,
		Email:  &body.Email,
		Status: &body.Status,
	})

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user read error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user read successfully",
		"user":    ReadUser,
	})
}

func (handler *BaseHandler) UpdateUser(context *fiber.Ctx) error {
	// Parse do corpo da solicitação para obter os dados de atualização
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
			"message": "Invalid request body",
			"errors":  errorMessages,
		})
	}

	updateParams := user.IUpdate{
		ID: body.ID,
	}

	if body.Email != nil {
		println("deu bom email")
		UpdateEmail := *body.Email
		handler.Service.UpdateEmail(body.ID, UpdateEmail)
	}

	if body.Status != nil {
		println("deu bom status")
		Updatestatus := *body.Status
		handler.Service.UpdateStatus(updateParams.ID, Updatestatus)
	}

	if body.Password != nil {
		println("deu bom senha")
		UpdatePass := *body.Password
		handler.Service.UpdatePassword(updateParams.ID, UpdatePass)
	}

	if body.Name != nil {
		println("deu bom nome")
		Updatename := *body.Name
		println(Updatename)
		handler.Service.UpdateName(updateParams.ID, Updatename)
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    updateParams,
	})
}

func (handler *BaseHandler) DeleteUser(context *fiber.Ctx) error {
	userIDstr := context.Params("id")

	userIDint, err := strconv.Atoi(userIDstr)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})
	}

	if userIDint < 0 {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	userID := uint(userIDint)

	deletedUser, serviceError := handler.Service.DeleteUser(userID)
	if serviceError != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting user",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user deleted successfully",
		"user":    deletedUser,
	})
}

func (handler *BaseHandler) SortProjects(context *fiber.Ctx) error {

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

	ProjectArrays, serviceError := handler.Service.SortProjects(body.ID)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user read error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Projects sorted by user succresfully",
		"array":   ProjectArrays,
	})
}
