package task

import (
	"orb-api/models"
	"orb-api/repositories/task"
)

type (
	Service struct {
		TaskRepo *task.Repository
	}

	Interface interface {
		AssignTask(idTask uint, idUser uint) (*models.Task, error)
	}
)