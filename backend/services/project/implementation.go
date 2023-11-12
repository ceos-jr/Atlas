package project

import (
	"orb-api/models"
	"orb-api/repositories/project"
)

func SetupProjectService(repository *project.Repository) *Service {
	return &Service{
		ProjectRepo: repository,
	}
}

func (service *Service) CreateProject(createData project.ICreate) (*models.Project, error) {
	NewProject, err := service.ProjectRepo.Create(createData)

	if err != nil {
		return nil, err
	}

	return NewProject, nil
}
