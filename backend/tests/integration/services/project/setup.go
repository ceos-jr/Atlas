package projectservicetest

import (
	"orb-api/config"
	"orb-api/models"
	"orb-api/services/project"
	
	repository "orb-api/repositories"

	"github.com/stretchr/testify/suite"
)

type TestSuit struct {
	suite.Suite
	Repo        *repository.Repository
	ProjectService *project.Service
	MockProjects []models.Project
}

// SetupSuite Executed before all tests
func (suite *TestSuit) SetupSuite() {
	repository, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repository
  	suite.ProjectService = project.SetupProjectService(&repository.Project)
}