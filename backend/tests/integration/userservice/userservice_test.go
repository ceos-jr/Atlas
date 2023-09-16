package userservicetest

import (
	userservice "orb-api/services/user"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *UserServiceTestSuit) TestCreateNewUser() {
	newUser, createErr := suite.userservice.CreateNewUser((userservice.ICreateUser{
		Name:     "exampleexample",
		Email:    "example@example.com.br",
		Password: suite.MockUsers[0].Password,
	}))

	suite.Nil(createErr, "Create error must be nil")
	suite.Equal("exampleexample", newUser.Name)
	suite.Equal("example@example.com.br", newUser.Email)
	suite.Equal(true, userservice.PasswordMatch(suite.MockUsers[0].Password, newUser.Password))

	suite.MockUsers[1] = *newUser
}

func (suite *UserServiceTestSuit) TestCreateNewUserErr() {
	invalidName := ""
	invalidEmail := ""
	invalidPassword := ""

	_, createErr := suite.userservice.CreateNewUser((userservice.ICreateUser{
		Name:     invalidName,
		Email:    "example@example.com",
		Password: suite.MockUsers[0].Password,
	}))

	suite.Equal("Invalid name", createErr.Error(), "Expected to have an error")

	_, createErr = suite.userservice.CreateNewUser((userservice.ICreateUser{
		Name:     "exampleexample2",
		Email:    invalidEmail,
		Password: suite.MockUsers[0].Password,
	}))

	suite.Equal("Invalid email", createErr.Error(), "Expected to have an error")

	_, createErr = suite.userservice.CreateNewUser((userservice.ICreateUser{
		Name:     "exampleexample3",
		Email:    "example@example.com.net",
		Password: invalidPassword,
	}))

	suite.Equal("Invalid password size", createErr.Error(), "Expected to have an error")

	_, createErr = suite.userservice.CreateNewUser((userservice.ICreateUser{
		Name:     "exampleexample3",
		Email:    suite.MockUsers[0].Email,
		Password: suite.MockUsers[0].Password,
	}))

	suite.Equal("This email is already being used", createErr.Error(), "Expected to have an error")

	_, createErr = suite.userservice.CreateNewUser((userservice.ICreateUser{
		Name:     suite.MockUsers[0].Name,
		Email:    "example@example.com.net",
		Password: suite.MockUsers[0].Password,
	}))

	suite.Equal("This username is already being used", createErr.Error(), "Expected to have an error")
}

func (suite *UserServiceTestSuit) TestUpdateName() {
	newName := "newName"
	updateName, updateErr := suite.userservice.UpdateName(suite.MockUsers[0].ID, newName)

	suite.Nil(updateErr, "Update error must be nil")
	suite.Equal(updateName.Name, newName)
}

func (suite *UserServiceTestSuit) TestUpdateNameErr() {
	invalidID := uint(1231313)
	invalidName := ""

	_, updateErr := suite.userservice.UpdateName(invalidID, "randomrandomName")
	suite.Equal("Invalid user id", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.userservice.UpdateName(suite.MockUsers[0].ID, invalidName)
	suite.Equal("Invalid username size", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.userservice.UpdateName(suite.MockUsers[0].ID, suite.MockUsers[1].Name)
	suite.Equal("This name is already being used", updateErr.Error(), "Expected to have an error")

}

func (suite *UserServiceTestSuit) TestUpdatePassword() {
	newPassword := "newestPassword"
	updatePassword, updateErr := suite.userservice.UpdatePassword(suite.MockUsers[0].ID, newPassword)

	suite.Nil(updateErr, "Update error must be nil")
	suite.Equal(true, userservice.PasswordMatch(newPassword, updatePassword.Password))
}

func (suite *UserServiceTestSuit) TestUpdatePasswordErr() {
	invalidID := uint(1231313)
	invalidPassword := ""

	_, updateErr := suite.userservice.UpdatePassword(invalidID, "randomrandomPassword")
	suite.Equal("Invalid user id", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.userservice.UpdatePassword(suite.MockUsers[0].ID, invalidPassword)
	suite.Equal("Invalid password size", updateErr.Error(), "Expected to have an error")
}

func (suite *UserServiceTestSuit) TestUpdateEmail() {
	newEmail := "newemail@example.com"
	updateEmail, updateErr := suite.userservice.UpdateEmail(suite.MockUsers[0].ID, newEmail)

	suite.Nil(updateErr, "Update erros must be nil")
	suite.Equal(updateEmail.Email, newEmail)
}

func (suite *UserServiceTestSuit) TestUpdateEmailErr() {
	invalidID := uint(1231313)
	invalidEmail := ""

	_, updateErr := suite.userservice.UpdateEmail(invalidID, "randomemail@example.com")
	suite.Equal("Invalid user id", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.userservice.UpdateEmail(suite.MockUsers[0].ID, invalidEmail)
	suite.Equal("Invalid email size", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.userservice.UpdateEmail(suite.MockUsers[0].ID, suite.MockUsers[1].Email)
	suite.Equal("This email is already being used", updateErr.Error(), "Expected to have an error")
}

func (suite *UserServiceTestSuit) TestUpdateStatus() {
	newStatus := uint(2)
	updateStatus, updateErr := suite.userservice.UpdateStatus(suite.MockUsers[0].ID, newStatus)

	suite.Nil(updateErr, "Update erros must be nil")
	suite.Equal(updateStatus.Status, newStatus)
}

func (suite *UserServiceTestSuit) TestUpdateStatusErr() {
	invalidID := uint(1231313)
	invalidStatus := uint(11)

	_, updateErr := suite.userservice.UpdateStatus(invalidID, uint(2))
	suite.Equal("Invalid user id", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.userservice.UpdateStatus(suite.MockUsers[0].ID, invalidStatus)
	suite.Equal("Invalid status", updateErr.Error(), "Expected to have an error")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuit))
}
