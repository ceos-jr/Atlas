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


func (service *Service) OrganizeTasks(UserId uint) (*[]models.Task, error) {
	if !service.TaskRepo.ValidUser(UserId){
		return nil, errors.New("invalid user id")
	}

	taskArray, readErr := service.TaskRepo.ReadBy(task.IReadBy{
		AssignedTo: &UserId,
	})


	if readErr != nil {
		return nil, readErr
	}

	if len(taskArray) == 0 {
		return nil, errors.New("no tasks to organize")
	}
	//use of the insertion sort to organize the array
	n := len(taskArray)
    for i := 1; i < n; i++ {
        key := taskArray[i]
        j := i - 1
   
        for j >= 0 && taskArray[j].Deadline.Unix() > key.Deadline.Unix() {
            taskArray[j+1] = taskArray[j]
            j = j - 1
		}
        taskArray[j+1] = key
	}

	return &taskArray, nil
	
	
	
}
