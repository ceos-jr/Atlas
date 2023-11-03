package task

import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/task"
)

func SetupTaskService(repository *task.Repository) *Service {
	return &Service{
		TaskRepo: repository,
	}
}

func (service *Service) AssignTask(idTask uint, idUser uint) (*models.Task, error) {
	if !service.TaskRepo.ValidUser(idUser) {
		return nil, errors.New("invalid user")
	}

	updateAssign, updateErr := service.TaskRepo.Update(task.IUpdate{
		ID:         idTask,
		AssignedTo: &idUser,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	return updateAssign, nil
}
