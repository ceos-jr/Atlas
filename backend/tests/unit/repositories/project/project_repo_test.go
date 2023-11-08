package projectrepotest

import (
	"orb-api/repositories/project"
	"testing"
	"github.com/stretchr/testify/suite"
)

func (suite *ProjectTestSuite) TestCreateProject() {
	project, createErr := suite.Repo.Project.Create(project.ICreate{
		Name: "Projeto",
		AdmID: suite.MockUser[0].ID,
		Sector: 1,
	})

	suite.Nil(createErr, "Create error must be nil")
	suite.Equal(suite.MockProject[0].Name, project.Name, "Name does not match")
	suite.Equal(suite.MockProject[0].AdmID, project.AdmID, "Adm ID does not match")
	suite.Equal(suite.MockProject[0].Sector, project.Sector, "Sector does not match")

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

	suite.Equal("no fields to read", readErr.Error(),
		"Empty fields it should return an error",
	)
}
func (suite *ProjectTestSuite) TestUpdateProject(){
	name := "Project"
	admid := uint(2)
	sector := uint(2)

	updatedProject, updateError := suite.Repo.Project.Update(project.IUpdate{
		ID:     suite.MockProject[0].ID,
		Name:   &name,
		AdmID:  &admid,
		Sector: &sector,
	})

	suite.Nil(updateError, "Update error must be nil")
	suite.Equal(updatedProject.Name, name,
		"Names do not match",
	)
	suite.Equal(updatedProject.Sector, sector,
		"Sector do not match",
	)

	suite.Equal(updatedProject.AdmID, admid,
		"Adm do not match",
	)
}

func GenerateString(length int) string {
	var generatedString = ""

	for i := 0; i < length; i++ {
		generatedString += "L"
	}
	return generatedString
}
func (suite *ProjectTestSuite) TestUpdateProjectErr() {
	invalidProjectName := GenerateString(129)
	invalidSector := uint(22)
	// invalidAdmId := suite.MockUser[0].ID

	// Test 01: Try to update with no fields
	_, updateError := suite.Repo.Project.Update(project.IUpdate{
		ID: suite.MockProject[0].ID,
	})

	suite.Equal("No fields to update", updateError.Error(),
	"Empty fields it should return an error",
	)

	// Test 02: Try to update with invalid Project Name
	_, updateError = suite.Repo.Project.Update(project.IUpdate{
		ID: suite.MockProject[0].ID,
		Name: &invalidProjectName,		      
	})
	
	suite.Equal("Invalid project name", updateError.Error(),
		"Invalid project name it should return an error",
	)

	// Test 03: Try to update with invalid sector
	_, updateError = suite.Repo.Project.Update(project.IUpdate{
		ID:   suite.MockProject[0].ID,
		Sector: &invalidSector,
	})

	suite.Equal("Invalid Sector", updateError.Error(),
		"Invalid sector it should return an error",
	)

	// Test 04: Try to update with invalid AdmId
	//_, updateError = suite.Repo.Project.Update(project.IUpdate{
	//	ID: suite.MockProject[0].ID,
	//	AdmID: &invalidAdmId,
	//})

	//suite.Equal("Invalid Adm ID", updateError.Error(),
	//	"Invalid Adm ID should return an error",
	//)
}

func (suite *ProjectTestSuite) TestDeleteProject() {
	newProject, _ := suite.Repo.Project.Create(project.ICreate{
		Name: "Atlas",
		Sector: 1,
		AdmID: suite.MockUser[0].ID,
	})

	deletedProject, deleteErr := suite.Repo.Project.Delete(project.IDelete{
		ID: newProject.ID,
	})

	suite.Nil(deleteErr, "Delete error must be nil")
	suite.Equal(newProject.ID, deletedProject.ID, "Expected to have the same ID")
}

func TestProjectRepository(test *testing.T) {
	suite.Run(test, new(ProjectTestSuite))
}