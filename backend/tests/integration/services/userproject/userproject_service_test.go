package userprojectservicetest

import (
	"orb-api/models"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (suite* TestSuit) TestAssignUser() {
	var Userprojects = make([]models.UsersProject, 1)

	ProjectID := suite.MockProjects[0].ID

	UserID := suite.MockUsers[0].ID

	NewUserProject, CreateErr := suite.ProjectService.AssignUser(suite.MockProjects[0].ID, suite.MockUsers[0].ID)

	suite.Nil(CreateErr, "Create error must be nil")

	suite.Equal(ProjectID, NewUserProject.ProjectID,"Project ID does not match")

	suite.Equal(UserID, NewUserProject.UserID,"User ID does not match")

	Userprojects[0] = *NewUserProject
	suite.MockUserProjects = Userprojects
}

func TestProjectService(test *testing.T) {
	suite.Run(test, new(TestSuit))
}