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
	idParam := context.Params("id")

	var id uint
    if idParam != "" {
        idValue, err := strconv.ParseUint(idParam, 10, 64)
        if err != nil {
            return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "message": "Erro durante a leitura de usuários",
                "error":   "ID inválido",
            })
        }
        id = uint(idValue)
    }

    readUserParams := user.IReadBy{
		ID: &id,
    }

    usersArray, serviceError := handler.Service.ReadUser(readUserParams)

    if serviceError != nil {
        return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Erro durante a leitura de usuários",
            "error":   serviceError.Error(),
        })
    }
    return context.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Leitura de usuários bem-sucedida",
        "users":   usersArray,
    })
}


func (handler *BaseHandler) ReadAllUsers(context *fiber.Ctx) error {
	usersArray, serviceError := handler.Service.ReadUser(user.IReadBy{})

	if serviceError != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro durante a leitura de usuários",
			"error":   serviceError.Error(),
		})
	}
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Leitura de todos os usuários bem-sucedida",
		"users":   usersArray,
	})
}

func (handler *BaseHandler) UpdateUser(context *fiber.Ctx) error {
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
		UpdateEmail := *body.Email
		_, err := handler.Service.UpdateEmail(body.ID, UpdateEmail)
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao atualizar o e-mail",
				"error":   err.Error(),
			})
		}
	}

	if body.Status != nil {
		Updatestatus := *body.Status
		_, err := handler.Service.UpdateStatus(updateParams.ID, Updatestatus)
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao atualizar o status",
				"error":   err.Error(),
			})
		}
	}

	if body.Password != nil {
		println("Entrou senha")
		UpdatePass := *body.Password
		_, err := handler.Service.UpdatePassword(updateParams.ID, UpdatePass)
		if err != nil {
			println("Deu merda na senha")
			println(err.Error())
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao atualizar a senha",
				"error":   err.Error(),
			})
		}
	}

	if body.Name != nil {
		Updatename := *body.Name
		_, err := handler.Service.UpdateName(updateParams.ID, Updatename)
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao atualizar o nome",
				"error":   err.Error(),
			})
		}
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
