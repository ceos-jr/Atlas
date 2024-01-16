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
	userRole, createErr := suite.Repo.UserRole.Create(userrole.ICreate{
		RoleID: suite.MockRole[1].ID,
		UserID: suite.MockUser[1].ID,
	})

	suite.Nil(createErr)
	suite.Equal(userRole.RoleID, suite.MockRole[1].ID, UnmatchedRoleID)
	suite.Equal(userRole.UserID, suite.MockUser[1].ID, UnmatchedUserID)

	suite.MockUserRole[1] = *userRole
}

func (suite *UserRoleRepoTestSuite) TestReadAllUserRoles() {
	userRoles, readErr := suite.Repo.UserRole.ReadAll()

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(2, len(userRoles), "Expected to have two user roles")
	suite.Equal(suite.MockUserRole[0].ID, userRoles[0].ID,
		"Expected to have the same id 0")
	suite.Equal(suite.MockUserRole[1].ID, userRoles[1].ID,
		"Expected to have the same id 1")
}

func (suite *UserRoleRepoTestSuite) TestReadByRoleIDUserRoles() {
	userRoles, readErr := suite.Repo.UserRole.ReadBy(userrole.IReadBy{
		RoleID: &suite.MockRole[0].ID,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(suite.MockUserRole[0].RoleID, userRoles[0].RoleID, "Expected to have same ID")
	suite.Equal(suite.MockUserRole[0].UserID, userRoles[0].UserID, "Expected to have same ID")
}

func (suite *UserRoleRepoTestSuite) TestReadByUserIDUserRoles() {
	userRoles, readErr := suite.Repo.UserRole.ReadBy(userrole.IReadBy{
		UserID: &suite.MockUser[0].ID,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(suite.MockUserRole[0].RoleID, userRoles[0].RoleID, "Expected to have same ID")
	suite.Equal(suite.MockUserRole[0].UserID, userRoles[0].UserID, "Expected to have same ID")
}

func (suite *UserRoleRepoTestSuite) TestReadByErr() {
	_, readErr := suite.Repo.UserRole.ReadBy(userrole.IReadBy{})

	suite.NotNil(readErr, "Read Error shouldn't be Nil (no fields to read)")
}

func (suite *UserRoleRepoTestSuite) TestUpdateUserRole() {
	ID := suite.MockUserRole[0].ID

	userRole, updateErr := suite.Repo.UserRole.Update(userrole.IUpdate{
		UserRoleID: ID,
		RoleID:     &suite.MockRole[1].ID,
	})

	suite.Nil(updateErr, "Update error should be nil")
	suite.Equal(suite.MockUserRole[0].ID, userRole.ID, "Should have been updated")
	suite.Equal(suite.MockRole[1].ID, userRole.RoleID, "Should have been updated")
}

func (suite *UserRoleRepoTestSuite) TestUpdateUserRoleErr() {
	invalidUserRoleID := 0
	var invalidUserID uint = 50
	var invalidRoleID uint = 1234

	// 1. Tentar atualizar sem campos preenchidos
	_, updateErr := suite.Repo.UserRole.Update(userrole.IUpdate{
		UserRoleID: suite.MockUserRole[0].ID,
	})

	suite.NotNil(updateErr, "Update Error shouldn't be Nil (no fields to update)")

	// 2. Tentar atualizar com UserRoleID inválido

	_, updateErr = suite.Repo.UserRole.Update(userrole.IUpdate{
		UserRoleID: uint(invalidUserRoleID),
		RoleID:     &suite.MockRole[0].ID,
	})

	suite.NotNil(updateErr, "Update Error shouldn't be Nil (no fields to update)")

	// 3. Tentar atualizar com RoleID inválido

	_, updateErr = suite.Repo.UserRole.Update(userrole.IUpdate{
		UserRoleID: suite.MockUserRole[0].ID,
		RoleID:     &invalidRoleID,
	})

	suite.NotNil(updateErr, "Update Error shouldn't be Nil (invalid RoleID)")

	// 4. Tentar atualizar com UserID inválido

	_, updateErr = suite.Repo.UserRole.Update(userrole.IUpdate{
		UserRoleID: suite.MockUserRole[0].ID,
		UserID:     &invalidUserID,
	})

	suite.NotNil(updateErr, "Update Error shouldn't be Nil (invalid UserID)")

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
