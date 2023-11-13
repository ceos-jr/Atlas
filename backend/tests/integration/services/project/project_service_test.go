package projectservicetest

import (
	"orb-api/models"
	"testing"
	"fmt"
	
	"github.com/stretchr/testify/suite"
)

func (suite *TestSuit) TestCreateNewProject () {
	var projects = make([]models.Project, 1)

	project, createErr := suite.ProjectService.CreateProject("Atlas",1,1)

	fmt.Print("\n",project.Sector)
	fmt.Print("\n",project.AdmID)

	suite.Nil(createErr, "Create error must be nil")

	suite.Equal("Atlas", project.Name, "Name does not match")
	suite.Equal(uint(1), project.Sector, "Sector does not match")
	suite.Equal(uint(1), project.AdmID, "Admin ID does not match")

	projects[0] = *project
	suite.MockProjects = projects
}

func TestProjectService(test *testing.T) {
	suite.Run(test, new(TestSuit))
}