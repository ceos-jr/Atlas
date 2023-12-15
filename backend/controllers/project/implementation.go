package project

import (
	"github.com/gofiber/fiber/v2"
)

type CreateProjectRequestBody struct {
	Name 		string `json:"name" validate:"required"`
	SectorID	uint`json:"sectorid" validate:"required"`
	AdmID		uint`json:"admid" validate:"required"`
}

type CreateUserProjectRequestBody struct {
	ProjectID	uint`json:"projectid" validate:"required"`
	UserID	uint`json:"userid" validate:"required"`
}

type CreateTaskProjectRequestBody struct {
	ProjectID	uint`json:"projectid" validate:"required"`
	TaskID	uint`json:"taskid" validate:"required"`
}

type CreateSortTaskRequestBody struct {
	ProjectID	uint`json:"projectid" validate:"required"`
}

func (handler *BaseHandler) CreateProject(context *fiber.Ctx) error{
	body := new(CreateProjectRequestBody)

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
		"project":    newProject,
	})
}

func (handler *BaseHandler) AssignUser(context *fiber.Ctx) error{
	body := new(CreateUserProjectRequestBody)

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

	userassign, serviceError := handler.Service.AssignUser(
		body.ProjectID,
		body.UserID,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error while assigning user",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user assigned successfully",
		"project":    userassign,
	})
}

func (handler *BaseHandler) AssignTask(context *fiber.Ctx) error{
	body := new(CreateUserProjectRequestBody)

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

	taskassign, serviceError := handler.Service.AssignTask(
		body.ProjectID,
		body.UserID,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error while assigning task",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "task assigned successfully",
		"project":    taskassign,
	})
}

func (handler *BaseHandler) SortByDeadline(context *fiber.Ctx) error{
	body := new(CreateUserProjectRequestBody)

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

	sortedtask, serviceError := handler.Service.SortByDeadline(
		body.ProjectID,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while sorting tasks",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "tasks found and sorted successfully",
		"project":    sortedtask,
	})
}
