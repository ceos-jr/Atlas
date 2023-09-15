package userolerepotest

import (
	userrole "orb-api/repositories/user_role"
)

const (
	UnmatchedRoleID = "RoleID and Role's ID don't match."
	UnmatchedUserID = "UserID and User's ID don't match."
)

func (suite *UserRoleRepoTestSuite) TestCreateUserRole() {
	user_role, createErr := suite.Repo.UserRole.Create(userrole.ICreate{
		RoleID: suite.MockRole[0].ID,
		UserID: suite.MockUser[0].ID,
	})

	suite.Nil(createErr)
	suite.Equal(user_role.RoleID, suite.MockRole[0].ID, UnmatchedRoleID)
	suite.Equal(user_role.UserID, suite.MockUser[0].ID, UnmatchedUserID)

}

func (suite *UserRoleRepoTestSuite) TestReadByUserRoles() {
	user_role, readErr := suite.Repo.UserRole.ReadBy(userrole.IReadBy{
		RoleID: &suite.MockRole[0].ID,
		UserID: &suite.MockUser[0].ID,
	})

}
