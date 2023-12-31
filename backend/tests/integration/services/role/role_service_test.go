package roleservicetest

import (
	"orb-api/models"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *TestSuite) TestCreateRole() {
	var roles = make([]models.Role, 1)

	role, createErr := suite.Service.NewRole("Role01", "This is a test")
	suite.Nil(createErr, "Create error must be nil")
	suite.Equal("This is a test", role.Description, "Description does not match")
	suite.Equal("Role01", role.Name, "Name does not match")

	roles[0] = *role
	suite.MockRoles = roles

}

func (suite *TestSuite) TestUpdateName() {
	role, updateErr := suite.Service.UpdateName(suite.MockRoles[0].ID, "NewRole")
	suite.Nil(updateErr, "Update error must be nil")
	suite.Equal("NewRole", role.Name, "Name does not match")
}

func (suite *TestSuite) TestUpdateDescription() {
	role, updateErr := suite.Service.UpdateDescription(suite.MockRoles[0].ID, "New Description")
	suite.Nil(updateErr, "Update error must be nil")
	suite.Equal("New Description", role.Description, "Description does not match")
}

func TestRoleRepository(test *testing.T) {
	suite.Run(test, new(TestSuite))
}
