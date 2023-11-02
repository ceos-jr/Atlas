package task

import (
	"orb-api/models"
	"orb-api/repositories/task"
)

func SetupTaskService(repository *task.Repository) *Service {
	return &Service{
		TaskRepo: repository,
	}
}

func (service *Service) MarkTaskAsCompleted(id uint) (*models.Task, error) {
	status := uint(1)
	statusp := &status

	taskUpdate, updateErr := service.TaskRepo.Update(task.IUpdate{
		ID:     id,
		Status: statusp,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	return taskUpdate, nil
}
