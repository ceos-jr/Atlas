package task

import (
	"errors"
	"orb-api/models"
	"orb-api/repositories/task"
)
type taskList []models.Task

func (t taskList) less(i int,j int) bool {
	return (t[i].Deadline).Before(t[j].Deadline)
}

func (t taskList) swap(i int,j int) {
	t[i], t[j] = t[j], t[i]
}

func (t taskList) Len() int {
	return len(t)
}

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

func (service *Service) AssignTask(idTask uint, idUser uint) (*models.Task, error) {
	updateAssign, updateErr := service.TaskRepo.Update(task.IUpdate{
		ID:         idTask,
		AssignedTo: &idUser,
	})

	if updateErr != nil {
		return nil, updateErr
	}

	return updateAssign, nil
}


func (service *Service) OrganizeTasks(idUser uint) (*models.Task, error) {
	var taskSlice []taskList

	taskArray, readErr := service.TaskRepo.ReadBy(task.IReadBy{
		AssignedTo: &idUser,
	})

	if readErr != nil {
		return nil, readErr
	}

	if len(taskArray) == 0 {
		return nil, errors.New("No tasks to organize")
	}
	
	
	
	
}
