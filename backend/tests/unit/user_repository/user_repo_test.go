package userrepotest

import (
	"orb-api/repositories/user"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (suite *UserRepoTestSuite) TestCreateUser() {
	user, createErr := suite.Repo.User.Create(user.ICreate{
		Name:    "User 01",
		Email:  "user01@example.com",
		Password: "123456",
		Status:  1,
})

	suite.Nil(createErr, "Create error must be nil")
	suite.Equal("User 01", user.Name, "Name does not match")
	suite.Equal("user01@example.com", user.Email, "Email does not match")
	suite.Equal(uint(1), user.Status, "Status does not match")
	suite.Equal("123456", user.Password, "Password does not match")

	suite.MockUsers[0] = *user
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
