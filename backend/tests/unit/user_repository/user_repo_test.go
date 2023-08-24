package userrepotest

import (
	"orb-api/repositories/user"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *UserRepoTestSuite) TestCreateUser() {
	user, createErr := suite.Repo.User.Create(user.ICreate{
		Name:     "User 01",
		Email:    "user01@example.com",
		Password: "12345678",
		Status:   1,
	})

	suite.Nil(createErr, "Create error must be nil")
	suite.Equal("User 01", user.Name, "Name does not match")
	suite.Equal("user01@example.com", user.Email, "Email does not match")
	suite.Equal(uint(1), user.Status, "Status does not match")
	suite.Equal("12345678", user.Password, "Password does not match")

	suite.MockUsers = append(suite.MockUsers, *user)
}

func (suite *UserRepoTestSuite) TestReadAllUsers() {
	users, readErr := suite.Repo.User.ReadAll(user.IReadAll{})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(1, len(users), "Expected to have one user")
	suite.Equal(suite.MockUsers[0].ID, users[0].ID, "Expected to have the same ID")
}

func (suite *UserRepoTestSuite) TestReadUserByID() {
	users, readErr := suite.Repo.User.ReadBy(user.IReadBy{
		ID: &suite.MockUsers[0].ID,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(1, len(users), "Expected to have one user")
	suite.Equal(suite.MockUsers[0].ID, users[0].ID, "Expected to have the same ID")
}

func (suite *UserRepoTestSuite) TestReadUserByName() {
	users, readErr := suite.Repo.User.ReadBy(user.IReadBy{
		Name: &suite.MockUsers[0].Name,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(1, len(users), "Expected to have one user")
	suite.Equal(suite.MockUsers[0].ID, users[0].ID, "Expected to have the same ID")

}

func (suite *UserRepoTestSuite) TestReadUserByEmail() {
	users, readErr := suite.Repo.User.ReadBy(user.IReadBy{
		Email: &suite.MockUsers[0].Email,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(1, len(users), "Expected to have one user")
	suite.Equal(suite.MockUsers[0].ID, users[0].ID, "Expected to have the same ID")
}

func (suite *UserRepoTestSuite) TestReadUserByStatus() {
	users, readErr := suite.Repo.User.ReadBy(user.IReadBy{
		Status: &suite.MockUsers[0].Status,
	})

	suite.Nil(readErr, "Read error must be nil")
	suite.Equal(suite.MockUsers[0].ID, users[0].ID, "Expected to have the same ID")

}

func (suite *UserRepoTestSuite) TestReadByErr() {
	_, readErr := suite.Repo.User.ReadBy(user.IReadBy{})

	suite.Equal("No fields to read", readErr.Error(), "Expected to have fields error")

}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (suite *UserRepoTestSuite) TestUpdateUser() {
	name := "Iguinho"
	email := "iguinho@email.com"
	status := uint(3)

	updatedUser, updateError := suite.Repo.User.Update(user.IUpdate{
		ID:     suite.MockUsers[1].ID,
		Name:   &name,
		Email:  &email,
		Status: &status,
	})

	suite.Nil(updateError, "Update error must be nil")
	suite.Equal(name, updatedUser.Name,
		"Names do not match",
	)
	suite.Equal(email, updatedUser.Email,
		"Emails do not match",
	)
	suite.Equal(status, updatedUser.Status,
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
func (suite *UserRepoTestSuite) TestUpdateUserErr() {
	invalidName := GenerateString(129)
	invalidEmail := GenerateString(129)
	invalidStatus := uint(77)

	// Teste 1: Tentativa de atualizar sem campos
	_, updateError := suite.Repo.User.Update(user.IUpdate{
		ID: suite.MockUsers[1].ID,
	})

	suite.Equal("No fields to update", updateError.Error(),
		"Empty fields it should return an error",
	)

	// Teste 2: Tentativa de atualizar com status inválido
	_, updateError = suite.Repo.User.Update(user.IUpdate{
		ID:     suite.MockUsers[1].ID,
		Status: &invalidStatus,
	})

	suite.Equal("Invalid user status", updateError.Error(),
		"Invalid user status it should return an error",
	)

	// Teste 3: Tentativa de atualizar com nome inválido
	_, updateError = suite.Repo.User.Update(user.IUpdate{
		ID:   suite.MockUsers[1].ID,
		Name: &invalidName,
	})

	suite.Equal("Invalid name", updateError.Error(),
		"Invalid user name it should return an error",
	)

	// Teste 4: Tentativa de atualizar com email inválido
	_, updateError = suite.Repo.User.Update(user.IUpdate{
		ID:    suite.MockUsers[1].ID,
		Email: &invalidEmail,
	})

	suite.Equal("Invalid user passed to createBy", updateError.Error(),
		"Invalid email should return an error",
	)

}

func (suite *UserRepoTestSuite) TestDeleteUser() {
	newUser, _ := suite.Repo.User.Create(user.ICreate{
		Name:     "vanilla",
		Email:    "vanilla@email.com",
		Status:   1,
		Password: "123454",
	})

	deletedUser, deleteErr := suite.Repo.User.Delete(user.IDelete{
		ID: newUser.ID,
	})

	suite.Nil(deleteErr, "Delete error must be nil")
	suite.Equal(newUser.ID, deletedUser.ID, "Expected to have the same id")
}
