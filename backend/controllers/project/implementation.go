package project

import (
	"orb-api/services/project"

	"github.com/gofiber/fiber/v2"
)

type CreateProjectRequestBody struct {
	ProjectID	uint`json:"projectid" validate:"required"`
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

func (handler *BaseHandler) UpdateProject(context *fiber.Ctx) error{
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

	updateProject, serviceError := handler.Service.UpdateProject(project.Update{
		ID: body.ProjectID,
		Name: &body.Name,
		Sector: &body.SectorID,
		AdmID: &body.AdmID,
	})

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "project update error",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "project updated successfully",
		"project":    updateProject,
	})
}


func (handler *BaseHandler) ListProjectbyUser(context *fiber.Ctx) error{
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

	listedprojects, serviceError := handler.Service.ListProjectbyUser(
		body.UserID,
	)

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error while listing user",
			"error":   serviceError.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "listed projects successfully",
		"project":    listedprojects,
	})
}

func (handler *BaseHandler) readProjects(context *fiber.Ctx) error{
	projectArr, serviceErr := handler.Service.ListProject()

	if serviceError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error while listing projects",
			"error": serviceErr.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber,Map{
		"message": "Projects listed succesfully",
		"array": projectArr 
	})
}
