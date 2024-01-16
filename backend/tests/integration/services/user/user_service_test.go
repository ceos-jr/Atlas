package userservicetest

import (
	userrepo "orb-api/repositories/user"
	"orb-api/services/user"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *TestSuit) TestCreateUser() {
	newUser, createErr := suite.Service.CreateUser(
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

	_, createErr := suite.Service.CreateUser(
		invalidName,
		"example@example.com",
		suite.MockUsers[0].Password,
	)

	suite.Equal("Invalid name", createErr.Error(), "Expected to have an error")

	_, createErr = suite.Service.CreateUser(
		"exampleexample2",
		invalidEmail,
		suite.MockUsers[0].Password,
	)

	suite.Equal("Invalid email", createErr.Error(), "Expected to have an error")

	_, createErr = suite.Service.CreateUser(
		"exampleexample3",
		"example@example.com.net",
		invalidPassword,
	)

	suite.Equal("Invalid password size", createErr.Error(), "Expected to have an error")

	_, createErr = suite.Service.CreateUser(
		"exampleexample3",
		suite.MockUsers[0].Email,
		suite.MockUsers[0].Password,
	)

	suite.Equal(
		"This email is already being used",
		createErr.Error(),
		"Expected to have an error",
	)

	_, createErr = suite.Service.CreateUser(
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

func (suite *TestSuit) TestDeleteUser() {
	deletedUser, deleteErr := suite.Service.DeleteUser(suite.MockUsers[1].ID)

	suite.Nil(deleteErr, "Delete erros must be nil")
	suite.Equal(deletedUser.Status, uint(1))
}

func (suite *TestSuit) TestDeleteUserErr() {
	id := uint(546)

	_, deleteErr := suite.Service.DeleteUser(id)
	suite.Equal("Invalid user id", deleteErr.Error(), "expected to have an error")

	_, deleteErr2 := suite.Service.DeleteUser(suite.MockUsers[2].ID)
	suite.Equal("User already disabled", deleteErr2.Error(), "expected to have an error")

}

func (suite *TestSuit) TestSortProjects() {
	projects, readErr := suite.Service.SortProjects(suite.MockUsers[0].ID)

	suite.Nil(readErr, "Create error must be nil")
	suite.Equal(1, len(projects), "lenght must be 1")
	suite.Equal(projects[0].ID, suite.MockProject.ID)
	suite.Equal(projects[0].AdmID, suite.MockProject.AdmID)
	suite.Equal(projects[0].Sector, suite.MockProject.Sector)
}

func (suite *TestSuit) TestReadAllUser() {
	userArray, readErr := suite.Service.ReadUser(userrepo.IReadBy{
		ID:     nil,
		Name:   nil,
		Email:  nil,
		Status: nil,
		Limit:  nil,
	})


	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(len(userArray), len(suite.MockUsers))

}

func (suite *TestSuit) TestReadByUser() {
	name := "Gabrigas5"
	userArray, readErr := suite.Service.ReadUser(userrepo.IReadBy{
		Name: &name,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(len(userArray), 1)
	suite.Equal(userArray[0].ID, suite.MockUsers[4].ID)
}

func (suite *TestSuit) TestReadByUserErr () {
	id := uint(5000)

	userArray, readErr := suite.Service.ReadUser(userrepo.IReadBy{
		ID: &id,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(len(userArray), 0)
}

func TestTaskRepository(t *testing.T) {
	suite.Run(t, new(TestSuit))
}
