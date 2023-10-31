package task

import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/task"
)

func SetupService(repository *task.Repository) *Service {
	return &Service{
		TaskRepo: repository,
	}
}

func (Service *Service) ConcludedTask(id uint)(*models.Task, error) {
	
}