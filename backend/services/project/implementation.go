package project

import (
	"orb-api/models"
	"orb-api/repositories/project"
	"errors"
)

func SetupProjectService(repository *project.Repository) *Service {
	return &Service{
		ProjectRepo: repository,
	}
}

func (service *Service) CreateProject(name string, Sector uint, AdmID uint) (*models.Project, error) {

	NewProject, err := service.ProjectRepo.Create(project.ICreate{
		Name:	name,
		Sector: Sector,
		AdmID:	AdmID,
	})

	if err != nil {
		return nil, err
	}

	return NewProject, nil
}

func (service *Service) AssignUser(ProjectID uint, UserID uint) (*models.UsersProject, error) {
	if !service.ProjectRepo.ValidProject(ProjectID) {
		return nil, errors.New("invalid Project passed to AssignUser")
	}

	if !service.ProjectRepo.ValidUser(UserID) {
		return nil, errors.New("invalid User passed to AssignUser")
	}

	NewUserProject := &models.UsersProject{
		UserID:    UserID,
		ProjectID: ProjectID,
	}

	return NewUserProject, nil
}

