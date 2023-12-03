package project

import (
	"orb-api/models"
	"orb-api/repositories/project"
	"orb-api/repositories/userproject"
	"orb-api/repositories/taskproject"
)

func SetupProjectService(repository1 *project.Repository, repository2 *userproject.Repository, repository3 *taskproject.Repository) *Service {
	return &Service{
		ProjectRepo: repository1,
		UserProjectRepo: repository2,
		TaskProjectRepo: repository3,
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

func (service *Service) AssignTask(ProjectID uint, TaskID uint) (*models.UsersProject, error) {

	NewTaskProject, Err := service.TaskProjectRepo.Create(taskproject.ICreate{
		ProjectID:	ProjectID,
		TaskID:		TaskID,
	})

	if Err != nil {
		return nil, Err
	}

	return NewTaskProject, nil
}

func (service *Service) SortTaskDeadline(ProjectID uint) ([]*models.Task, error) {

	TaskProjects, Err := srvice.TaskProjectRepo.ReadBy(taskproject.IReadBy{
		ProjectID: ProjectID,
	})

	if Err != nil {
		return nil, Err
	}

	
}