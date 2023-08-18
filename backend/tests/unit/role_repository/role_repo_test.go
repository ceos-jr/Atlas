package rolerepotest

import (
	"orb-api/models"
	"orb-api/repositories/role"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *RoleRepoTestSuite) TestCreateRole() {
	var roles = make([]models.Role, 1)

	role, createErr := suite.Repo.Role.Create(role.ICreate{
		Name:        "Role 01",
		Description: "This is a test",
	})
	suite.Nil(createErr, "Create error must be nil")
	suite.Equal("This is a test", role.Description, "Description does not match")
	suite.Equal("Role 01", role.Name, "Name does not match")

	roles[0] = *role

	suite.MockRoles = roles
}

func (suite *RoleRepoTestSuite) TestReadAllRoles() {
	limit := 1
	roles, readErr := suite.Repo.Role.ReadAll(role.IReadAll{
		Limit: &limit,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(1, len(roles), "Expected to have one role")
	suite.Equal(suite.MockRoles[0].ID, roles[0].ID, "Expected to have the same ID")
}

func (suite *RoleRepoTestSuite) TestUpdateRole() {
	updateNameTest := "Role Test"
	updateDescriptionTest := "This is another test"

	updatedrole, updateErr := suite.Repo.Role.Update(role.IUpdate{
		RoleID:      suite.MockRoles[0].ID,
		Name:        &updateNameTest,
		Description: &updateDescriptionTest,
	})

	suite.Nil(updateErr, "Update error must be nil")
	suite.Equal(suite.MockRoles[0].ID, updatedrole.ID, "Expected to have the same ID")
	suite.Equal(updatedrole.Name, "Role Test", "Expected to have updated the same name")
	suite.Equal(updatedrole.Description, "This is another test", "Expected to have the same description")
}

func (suite *RoleRepoTestSuite) TestUpdateRoleErr() {
	_, updateErr := suite.Repo.Role.Update(role.IUpdate{
		RoleID: suite.MockRoles[0].ID,
	})
	suite.Equal("No fields to update", updateErr.Error(), "Expected to have fields error")

	updateNameTest := "Role Test"
	updateDescriptionTest := "This is another test"

	_, updateErr = suite.Repo.Role.Update(role.IUpdate{
		RoleID:      0,
		Name:        &updateNameTest,
		Description: &updateDescriptionTest,
	})
	suite.Equal("WHERE conditions required", updateErr.Error(), "Expected to have an ID error")

}

func (suite *RoleRepoTestSuite) TestReadByRole() {
	readByRole, readByErr := suite.Repo.Role.ReadBy(role.IReadBy{
		ID:          &suite.MockRoles[0].ID,
		Name:        &suite.MockRoles[0].Name,
		Description: &suite.MockRoles[0].Description,
	})

	suite.Nil(readByErr, "ReadBy error must be nil")

	readRole := readByRole[0]

	suite.Equal(suite.MockRoles[0].ID, readRole.ID, "Expected to hame same ID")
	suite.Equal(suite.MockRoles[0].Name, readRole.Name, "Expected to hame same name")
	suite.Equal(suite.MockRoles[0].Description, readRole.Description, "Expected to hame same description")
}

func (suite *RoleRepoTestSuite) TestDeleteRole() {
	newrole, _ := suite.Repo.Role.Create(role.ICreate{
		Name:        "Role 02",
		Description: "This is a another test",
	})

	deletedrole, deleteErr := suite.Repo.Role.Delete(role.IDelete{
		RoleID: newrole.ID,
	})

	suite.Nil(deleteErr, "Delete error must be nil")
	suite.Equal(newrole.ID, deletedrole.ID, "Expected to have the same ID")
}

func TestRoleRepository(test *testing.T) {
	suite.Run(test, new(RoleRepoTestSuite))
}
