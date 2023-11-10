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
	//validate project name and sector; check if the admin exists? Didn't do it because we don't have sector's model and I don't understand exactly how Admin works.

	newProject, err := service.ProjectRepo.Create(createData)
	if err!= nil {
		return nil, err
	}

	//create and associate employees to the project?
	for _, employeeID := range createData.EmployeeIDs {
		userProject := models.UserProject{
			UserID: employeeID,
			ProjectID: newProject.ID,
		}
	
	//We are supposed to associate users to projects, but we don't have anything like that in repository. Since the issue is not about changing repository, this would be an alternative? :

	result := service.ProjectRepo.AddUserToProject(userProject)
	if result.Error != nil {
		//comment: Handle the error?
		return nil, result.Error
	}

	return newProject, nil
}

