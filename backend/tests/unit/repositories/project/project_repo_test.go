package projectrepotest

import (
	"orb-api/repositories/project"
	"testing"
	"github.com/stretchr/testify/suite"
)

func (suite *ProjectTestSuite) TestCreateProject() {
	project, createErr := suite.Repo.Project.Create(project.ICreate{
		Name:     "Project 01",
		AdmID:	  suite.MockUser[0].ID,
		Sector:	  1,
	})

	suite.Nil(createErr, "Create error must be nil")
	suite.Equal("Project 01", project.Name, "Name does not match")
	suite.Equal(suite.MockProject[0],project.AdmID, "Adm ID does not match")
	suite.Equal(uint(1), project.Sector, "Sector does not match")

	suite.MockProject[1] = *project
}

func (suite *ProjectTestSuite) TestReadByID(){
	project, readErr := suite.Repo.Project.ReadBy(project.IReadBy{
		ID: &suite.MockProject[0].ID,
	})
	
	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(1, len(project), "Expected to have one user")
	suite.Equal(suite.MockProject[0].ID, project[0].ID, "Expected to have the same ID")
}

func (suite *ProjectTestSuite) TestReadByAdmID() {
	project, readErr := suite.Repo.Project.ReadBy(project.IReadBy{
		AdmID: &suite.MockProject[0].AdmID,
	})

	suite.Nil(readErr, "Read error should be empty")
	suite.Equal(suite.MockProject[0].AdmID, project[0].AdmID,
		"The Adm must match",
	)
}

func (suite *ProjectTestSuite) TestReadBySector(){
	project, readErr := suite.Repo.Project.ReadBy(project.IReadBy{
		Sector: &suite.MockProject[0].Sector,
	})

	suite.Nil(readErr, "Read error should be empty")
	suite.Equal(suite.MockProject[0].Sector, project[0].Sector,
		"The Sector must match",
	)
}

func (suite *ProjectTestSuite) TestReadByErr(){
	_, readErr := suite.Repo.Project.ReadBy(project.IReadBy{})

	suite.Equal("No fields to read", readErr.Error(),
		"Empty fields it should return an error",
	)
}

func TestProjectRepository(test *testing.T) {
	suite.Run(test, new(ProjectTestSuite))
}