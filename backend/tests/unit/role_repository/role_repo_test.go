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
	suite.Equal("Sucessor do Lucas Braide", role.Description, "Description does not match")
	suite.Equal("Lucas Braide Jr.", role.Name, "Name does not match")

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
