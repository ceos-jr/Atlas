package user_roleservicetest

import (
	user_rolerepo "orb-api/repositories/user_role"
	rolerepo "orb-api/repositories/role"
	userrepo "orb-api/repositories/user"
	"testing"

	"github.com/stretchr/testify/suite"
)
func (suite *TestSuite) TestAssigneRole(){
	UserRole, err := suite.Service.AssigneRole(suite.MockUsers[0].ID, suite.MockRoles[0].ID)

	suite.Nil(err, "Assigne Role error must be nil")
	suite.Equal(suite.MockUsers[0].ID, UserRole.UserID)
	suite.Equal(suite.MockRoles[0].ID, UserRole.RoleID)
}

func (suite *TestSuite) TestAssigneRoleErr() {
	Invalid_UserID := uint(500)
	Invalid_RoleID := uint(500)
	
	_, err := suite.Service.AssigneRole(suite.MockUsers[0].ID,Invalid_RoleID,)
	suite.Equal("invalid role id", err.Error(), "Expected to have an error")
	
	_, err = suite.Service.AssigneRole(Invalid_UserID, suite.MockRoles[0].ID,)
	suite.Equal("invalid user id", err.Error(), "Expected to have an error")
}

func (suite *TestSuite) TestUnassignRole() {
	//novo user role
	newUserRole, errs := suite.Service.UserRoleRepo.Create(user_rolerepo.ICreate{
		RoleID: suite.MockRoles[1].ID,
		UserID: suite.MockUsers[1].ID,
	})
	if errs != nil{
		panic(errs)
	}
	//usa o Unassign role
	UserRole, err := suite.Service.UnassignRole(newUserRole.UserID, newUserRole.RoleID)
	
	suite.Nil(err, "Unassigne Role error must be nil")
	suite.Equal(newUserRole.UserID, UserRole.UserID)
	suite.Equal(newUserRole.RoleID, UserRole.RoleID)

	//Deleta o mock de user e role
	_, deleteErr := suite.Service.UserRepo.Delete(userrepo.IDelete{
		ID: suite.MockUsers[1].ID,
	})
	
	if deleteErr != nil {
		panic(deleteErr)
	}
	
	_, deleteErr = suite.Service.RoleRepo.Delete(rolerepo.IDelete{
		RoleID: suite.MockRoles[1].ID,
	})
	
	if deleteErr != nil {
		panic(deleteErr)
	}
}	
func (suite *TestSuite) TestUnassignRoleErr() {
	invalid_user_id := uint(1000)
	invalid_role_id := uint(1000)

	_, err := suite.Service.UnassignRole(
		invalid_user_id,
		suite.MockRoles[0].ID,
	)	
	suite.Equal("invalid user id", err.Error(), "Expected to have an error")

	_, err = suite.Service.UnassignRole(
		suite.MockUsers[0].ID,
		invalid_role_id,
	)
	suite.Equal("invalid role id", err.Error(), "Expected to have an error")
}	

func TestUserRoleRepository(test *testing.T) {
	suite.Run(test, new(TestSuite))
}	
