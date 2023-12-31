package userrepotest

import (
	"github.com/stretchr/testify/suite"
	"orb-api/repositories/user"
	"testing"
)

func (suite *TestSuite) TestCreateUser() {
	newUser, createErr := suite.Repo.User.Create(user.ICreate{
		Name:     "User 01",
		Email:    "user01@example.com",
		Password: "12345678",
		Status:   1,
	})

	suite.Nil(createErr, "Create error must be nil")
	suite.Equal("User 01", newUser.Name, "Name does not match")
	suite.Equal("user01@example.com", newUser.Email, "Email does not match")
	suite.Equal(uint(1), newUser.Status, "Status does not match")
	suite.Equal("12345678", newUser.Password, "Password does not match")

	suite.MockUsers[1] = *newUser
}

func (suite *TestSuite) TestReadAllUsers() {
	users, readErr := suite.Repo.User.ReadAll(user.IReadAll{})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(2, len(users), "Expected to have two users")
	suite.Equal(suite.MockUsers[0].ID, users[0].ID, "Expected to have the same ID")

}


func (suite *TestSuite) TestReadUserByID() {

	invalidID := uint(777)

	users, readErr := suite.Repo.User.ReadBy(user.IReadBy{
		ID: &suite.MockUsers[0].ID,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(1, len(users), "Expected to have one user")
	suite.Equal(suite.MockUsers[0].ID, users[0].ID, "Expected to have the same ID")

	users, readErr = suite.Repo.User.ReadBy(user.IReadBy{
		ID: &invalidID,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(0, len(users), "Expected to have one user")
}

func (suite *TestSuite) TestReadUserByName() {
	users, readErr := suite.Repo.User.ReadBy(user.IReadBy{
		Name: &suite.MockUsers[0].Name,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(2, len(users), "Expected to have two users")
	suite.Equal(suite.MockUsers[0].ID, users[0].ID, "Expected to have the same ID")

}

func (suite *TestSuite) TestReadUserByEmail() {
	users, readErr := suite.Repo.User.ReadBy(user.IReadBy{
		Email: &suite.MockUsers[0].Email,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(2, len(users), "Expected to have two users")
	suite.Equal(suite.MockUsers[0].ID, users[0].ID, "Expected to have the same ID")
}

func (suite *TestSuite) TestReadUserByStatus() {
	users, readErr := suite.Repo.User.ReadBy(user.IReadBy{
		Status: &suite.MockUsers[0].Status,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(suite.MockUsers[0].ID, users[0].ID, "Expected to have the same ID")

}

func (suite *TestSuite) TestReadByErr() {
	_, readErr := suite.Repo.User.ReadBy(user.IReadBy{})

	suite.Equal("No fields to read", readErr.Error(), "Expected to have fields error")
}

func (suite *TestSuite) TestUpdateUser() {
	name := "Igor"
	email := "iguinho@email.com"
	status := uint(3)

	updatedUser, updateError := suite.Repo.User.Update(user.IUpdate{
		ID:     suite.MockUsers[0].ID,
		Name:   &name,
		Email:  &email,
		Status: &status,
	})

	suite.Nil(updateError, "Update error must be nil")
	suite.Equal(updatedUser.Name, name,
		"Names do not match",
	)
	suite.Equal(updatedUser.Email, email,
		"Emails do not match",
	)
	suite.Equal(updatedUser.Status, status,
		"Status do not match",
	)
}

func GenerateString(length int) string {
	var generatedString = ""

	for i := 0; i < length; i++ {
		generatedString += "L"
	}
	return generatedString
}

func (suite *TestSuite) TestUpdateUserErr() {
	invalidName := GenerateString(129)
	invalidEmail := GenerateString(129)
	invalidStatus := uint(77)
	invalidPassword := "short"

	// Test 01: Try to update with no fields
	_, updateError := suite.Repo.User.Update(user.IUpdate{
		ID: suite.MockUsers[0].ID,
	})

	suite.Equal("No fields to update", updateError.Error(),
		"Empty fields it should return an error",
	)

	// Test 02: Try to update with invalid status
	_, updateError = suite.Repo.User.Update(user.IUpdate{
		ID:     suite.MockUsers[0].ID,
		Status: &invalidStatus,
	})

	suite.Equal("Invalid status", updateError.Error(),
		"Invalid user status it should return an error",
	)

	// Test 03: Try to update with invalid name
	_, updateError = suite.Repo.User.Update(user.IUpdate{
		ID:   suite.MockUsers[0].ID,
		Name: &invalidName,
	})

	suite.Equal("Invalid name", updateError.Error(),
		"Invalid user name it should return an error",
	)

	// Test 04: Try to update with invalid email
	_, updateError = suite.Repo.User.Update(user.IUpdate{
		ID:    suite.MockUsers[0].ID,
		Email: &invalidEmail,
	})

	suite.Equal("Invalid email", updateError.Error(),
		"Invalid email should return an error",
	)

	// Test 05: Try to update with invalid password
	_, updateError = suite.Repo.User.Update(user.IUpdate{
		ID:       suite.MockUsers[0].ID,
		Password: &invalidPassword,
	})

	suite.Equal("Invalid password", updateError.Error(),
		"Invalid password should return an error",
	)
}

func (suite *TestSuite) TestDeleteUser() {
	newUser, _ := suite.Repo.User.Create(user.ICreate{
		Name:     "vanilla",
		Email:    "vanilla@email.com",
		Status:   1,
		Password: "12345678",
	})

	deletedUser, deleteErr := suite.Repo.User.Delete(user.IDelete{
		ID: newUser.ID,
	})

	suite.Nil(deleteErr, "Delete error must be nil")
	suite.Equal(newUser.ID, deletedUser.ID, "Expected to have the same id")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
