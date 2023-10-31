package task

import (
	"orb-api/models"
	"orb-api/repositories/task"
)

func SetupTask(repository *task.Repository) *Service {
	return &Service{
		TaskRepo: repository,
	}
}

func (service *Service) ConcludedTask(id uint)(*models.Task, error) {

	status := uint(1)

	taskUpdate, updateErr := service.TaskRepo.Update(task.IUpdate{
		ID:     id,
		Status: &status,
	})

	if updateErr != nil{
		return nil, updateErr
	}

	return taskUpdate, nil
}

