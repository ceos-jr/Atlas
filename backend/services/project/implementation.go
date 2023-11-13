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