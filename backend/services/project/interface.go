package project

import (
	"orb-api/models"
	"orb-api/repositories/project"
	"orb-api/repositories/userproject"
	"orb-api/repositories/taskproject"
	"orb-api/repositories/task"
)

type (
	Update struct{
		ID     uint
    	Name   *string
    	Sector *uint
    	AdmID  *uint
	}

	Service struct {
		ProjectRepo *project.Repository
		UserProjectRepo	*userproject.Repository
		TaskProjectRepo *taskproject.Repository
		TaskRepo 		*task.Repository
	}

	Interface interface {
		CreateProject(name string, Sector uint, AdmID uint) (*models.Project, error)
		AssignUser(ProjectID uint, UserID uint) (*models.UsersProject, error)
		AssignTask(ProjectID uint, TaskID uint) (*models.TasksProject, error)
		SortTaskByDeadline(ProjectID uint) ([]models.Task, error)
		UpdateProject(Update) (*models.Project, error)
		ListProject() ([]models.Project, error)
	}
)
