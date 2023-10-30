package task

import (
	"orb-api/models"
	"orb-api/repositories/task"
)

type (
	Service struct{
		TaskRepo* task.Repository
	}

	Interface interface{
		ConcludedTask(id uint)(*models.Task, error)
	}
)