package userservicetest

import (
	"orb-api/services/user"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *TestSuit) TestCreateNewUser() {
	newUser, createErr := suite.Service.NewUser(
		"exampleexample",
		"example@example.com.br",
		suite.MockUsers[0].Password,
	)

	suite.Nil(createErr, "Create error must be nil")
	suite.Equal("exampleexample", newUser.Name)
	suite.Equal("example@example.com.br", newUser.Email)
	suite.Equal(true, user.PasswordMatch(suite.MockUsers[0].Password, newUser.Password))

	suite.MockUsers[1] = *newUser
}

func (suite *TestSuit) TestCreateNewUserErr() {
	invalidName := ""
	invalidEmail := ""
	invalidPassword := ""

	_, createErr := suite.Service.NewUser(
		invalidName,
		"example@example.com",
		suite.MockUsers[0].Password,
	)

	suite.Equal("Invalid name", createErr.Error(), "Expected to have an error")

	_, createErr = suite.Service.NewUser(
		"exampleexample2",
		invalidEmail,
		suite.MockUsers[0].Password,
	)

	suite.Equal("Invalid email", createErr.Error(), "Expected to have an error")

	_, createErr = suite.Service.NewUser(
		"exampleexample3",
		"example@example.com.net",
		invalidPassword,
	)

	suite.Equal("Invalid password size", createErr.Error(), "Expected to have an error")

	_, createErr = suite.Service.NewUser(
		"exampleexample3",
		suite.MockUsers[0].Email,
		suite.MockUsers[0].Password,
	)

	suite.Equal(
		"This email is already being used",
		createErr.Error(),
		"Expected to have an error",
	)

	_, createErr = suite.Service.NewUser(
		suite.MockUsers[0].Name,
		"example@example.com.net",
		suite.MockUsers[0].Password,
	)

	suite.Equal(
		"This username is already being used",
		createErr.Error(),
		"Expected to have an error",
	)
}

func (suite *TestSuit) TestUpdateName() {
	newName := "newName"
	updateName, updateErr := suite.Service.UpdateName(suite.MockUsers[0].ID, newName)

	suite.Nil(updateErr, "Update error must be nil")
	suite.Equal(updateName.Name, newName)
}

func (suite *TestSuit) TestUpdateNameErr() {
	invalidID := uint(1231313)
	invalidName := ""

	_, updateErr := suite.Service.UpdateName(invalidID, "randomrandomName")
	suite.Equal("Invalid user id", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.Service.UpdateName(suite.MockUsers[0].ID, invalidName)
	suite.Equal("Invalid username size", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.Service.UpdateName(
		suite.MockUsers[0].ID,
		suite.MockUsers[1].Name,
	)
	suite.Equal(
		"This name is already being used",
		updateErr.Error(),
		"Expected to have an error",
	)

}

func (suite *TestSuit) TestUpdatePassword() {
	newPassword := "newestPassword"
	updatePassword, updateErr := suite.Service.UpdatePassword(
		suite.MockUsers[0].ID, newPassword,
	)

	suite.Nil(updateErr, "Update error must be nil")
	suite.Equal(true, user.PasswordMatch(newPassword, updatePassword.Password))
}

func (suite *TestSuit) TestUpdatePasswordErr() {
	invalidID := uint(1231313)
	invalidPassword := ""

	_, updateErr := suite.Service.UpdatePassword(invalidID, "randomrandomPassword")
	suite.Equal("Invalid user id", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.Service.UpdatePassword(suite.MockUsers[0].ID, invalidPassword)
	suite.Equal("Invalid password size", updateErr.Error(), "Expected to have an error")
}

func (suite *TestSuit) TestUpdateEmail() {
	newEmail := "newemail@example.com"
	updateEmail, updateErr := suite.Service.UpdateEmail(suite.MockUsers[0].ID, newEmail)

	suite.Nil(updateErr, "Update erros must be nil")
	suite.Equal(updateEmail.Email, newEmail)
}

func (suite *TestSuit) TestUpdateEmailErr() {
	invalidID := uint(1231313)
	invalidEmail := ""

	_, updateErr := suite.Service.UpdateEmail(invalidID, "randomemail@example.com")
	suite.Equal("Invalid user id", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.Service.UpdateEmail(suite.MockUsers[0].ID, invalidEmail)
	suite.Equal("Invalid email size", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.Service.UpdateEmail(
		suite.MockUsers[0].ID,
		suite.MockUsers[1].Email,
	)
	suite.Equal(
		"This email is already being used",
		updateErr.Error(),
		"Expected to have an error",
	)
}

func (suite *TestSuit) TestUpdateStatus() {
	newStatus := uint(2)
	updateStatus, updateErr := suite.Service.UpdateStatus(
		suite.MockUsers[0].ID, newStatus,
	)

	suite.Nil(updateErr, "Update erros must be nil")
	suite.Equal(updateStatus.Status, newStatus)
}

func (suite *TestSuit) TestUpdateStatusErr() {
	invalidID := uint(1231313)
	invalidStatus := uint(11)

	_, updateErr := suite.Service.UpdateStatus(invalidID, uint(2))
	suite.Equal("Invalid user id", updateErr.Error(), "Expected to have an error")

	_, updateErr = suite.Service.UpdateStatus(suite.MockUsers[0].ID, invalidStatus)
	suite.Equal("Invalid status", updateErr.Error(), "Expected to have an error")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(TestSuit))
}
