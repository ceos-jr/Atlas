package project

import (
	"orb-api/models"
	"orb-api/repositories/project"
	"orb-api/repositories/userproject"
	"orb-api/repositories/taskproject"
	"orb-api/repositories/task"
)

func SetupProjectService(repository1 *project.Repository, repository2 *userproject.Repository, repository3 *taskproject.Repository, repository4 *task.Repository) *Service {
	return &Service{
		ProjectRepo: repository1,
		UserProjectRepo: repository2,
		TaskProjectRepo: repository3,
		TaskRepo: repository4,
	}
}

func (service *Service) CreateProject(name string, Sector uint, AdmID uint) (*models.Project, error) {

	NewProject, Err := service.ProjectRepo.Create(project.ICreate{
		Name:	name,
		Sector: Sector,
		AdmID:	AdmID,
	})

	if Err != nil {
		return nil, Err
	}

	return NewProject, nil
}

func (service *Service) AssignUser(ProjectID uint, UserID uint) (*models.UsersProject, error) {

	NewUserProject, Err := service.UserProjectRepo.Create(userproject.ICreate{
		ProjectID:	ProjectID,
		UserID:		UserID,
	})

	if Err != nil {
		return nil, Err
	}

	return NewUserProject, nil
}

func (service *Service) AssignTask(ProjectID uint, TaskID uint) (*models.TasksProject, error) {

	NewTaskProject, Err := service.TaskProjectRepo.Create(taskproject.ICreate{
		ProjectID:	ProjectID,
		TaskID:		TaskID,
	})

	if Err != nil {
		return nil, Err
	}

	return NewTaskProject, nil
}

func (service *Service) SortByDeadline(ProjectID uint) ([]models.Task, error) {

	TaskProjects, ReadErr := service.TaskProjectRepo.ReadBy(taskproject.IReadBy{
		ProjectID: &ProjectID,
	})

	if ReadErr != nil {
		return nil, ReadErr
	}

	Tasks := []models.Task{}

	for i := 0; i < len(TaskProjects); i++{
		App, Err := service.TaskRepo.ReadBy(task.IReadBy{
			ID: &(TaskProjects[i].TaskID),
		})

		if Err != nil {
			return nil, Err
		}

		Tasks = append(Tasks, App...)
	}

	Tasks = service.TaskRepo.Sort(Tasks)

	return Tasks, nil
}
