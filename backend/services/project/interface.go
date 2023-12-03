package project

import (
	"orb-api/models"
	"orb-api/repositories/project"
	"orb-api/repositories/userproject"
	"orb-api/repositories/taskproject"
)

type (
	Service struct {
		ProjectRepo *project.Repository
		UserProjectRepo	*userproject.Repository
		TaskProjectRepo *taskproject.Repository
	}

	Interface interface {
		CreateProject(name string, Sector uint, AdmID uint) (*models.Project, error)
		AssignUser(ProjectID uint, UserID uint) (*models.UsersProject, error)
		AssignTask(ProjectID uint, TaskID uint) (*models.TasksProject, error)
	}
)
