package userolerepotest

import (
	userrole "orb-api/repositories/user_role"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	UnmatchedRoleID = "RoleID and Role's ID don't match."
	UnmatchedUserID = "UserID and User's ID don't match."
)

func (suite *UserRoleRepoTestSuite) TestCreateUserRole() {
	user_role, createErr := suite.Repo.UserRole.Create(userrole.ICreate{
		RoleID: suite.MockRole[1].ID,
		UserID: suite.MockUser[1].ID,
	})

	suite.Nil(createErr)
	suite.Equal(user_role.RoleID, suite.MockRole[1].ID, UnmatchedRoleID)
	suite.Equal(user_role.UserID, suite.MockUser[1].ID, UnmatchedUserID)

	suite.MockUserRole[1] = *user_role
}

func (suite *UserRoleRepoTestSuite) TestReadAllUserRoles() {
	user_roles, readErr := suite.Repo.UserRole.ReadAll()

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(2, len(user_roles), "Expected to have two user roles")
	suite.Equal(suite.MockUserRole[0].ID, user_roles[0].ID,
		"Expected to have the same id 0")
	suite.Equal(suite.MockUserRole[1].ID, user_roles[1].ID,
		"Expected to have the same id 1")
}

func (suite *UserRoleRepoTestSuite) TestReadByRoleIDUserRoles() {
	user_roles, readErr := suite.Repo.UserRole.ReadBy(userrole.IReadBy{
		RoleID: &suite.MockRole[0].ID,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(suite.MockUserRole[0].RoleID, user_roles[0].RoleID, "Expected to have same ID")
	suite.Equal(suite.MockUserRole[0].UserID, user_roles[0].UserID, "Expected to have same ID")
}

func (suite *UserRoleRepoTestSuite) TestReadByUserIDUserRoles() {
	user_roles, readErr := suite.Repo.UserRole.ReadBy(userrole.IReadBy{
		UserID: &suite.MockUser[0].ID,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(suite.MockUserRole[0].RoleID, user_roles[0].RoleID, "Expected to have same ID")
	suite.Equal(suite.MockUserRole[0].UserID, user_roles[0].UserID, "Expected to have same ID")
}

func (suite *UserRoleRepoTestSuite) TestUpdateUserRole() {
	ID := suite.MockUserRole[0].ID

	user_role, updateErr := suite.Repo.UserRole.Update(userrole.IUpdate{
		UserRoleID: ID,
		RoleID:     &suite.MockRole[1].ID,
	})

	suite.Nil(updateErr, "Update error should be nil")
	suite.Equal(suite.MockUserRole[0].ID, user_role.ID, "Should have been updated")
	suite.Equal(suite.MockRole[1].ID, user_role.RoleID, "Should have been updated")
}

func (suite *UserRoleRepoTestSuite) TestDeleteUserRole() {
	newUserRole, _ := suite.Repo.UserRole.Create(userrole.ICreate{
		RoleID: suite.MockRole[1].ID,
		UserID: suite.MockUser[2].ID,
	})

	deletedUserRole, deleteErr := suite.Repo.UserRole.Delete(userrole.IDelete{
		UserRoleID: newUserRole.ID,
	})

	suite.Nil(deleteErr, "Delete error must be nil")
	suite.Equal(newUserRole.ID, deletedUserRole.ID, "Expected to have the same id")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRoleRepoTestSuite))
}
