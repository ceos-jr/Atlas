package userrepotest

import (
	"fmt"
	"orb-api/config"
	"orb-api/models"
	"orb-api/repositories"
	"orb-api/repositories/user"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	Repo      *repository.Repository
	MockUsers []models.User
}

// Executed before all tests
func (suite *UserRepoTestSuite) SetupSuite() {
	repo, setupError := config.SetupDB("../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
	suite.MockUsers = make([]models.User, 2)
	suite.SetupMocks()
}

// Executed after all tests
func (suite *UserRepoTestSuite) TearDownSuite() {
	for _, user := range suite.MockUsers {
		_, deleteErr := suite.Repo.User.Delete(user.IDelete{
			ID: user.ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}

func (suite *UserRepoTestSuite) SetupMocks() {
	for index := 0; index < 2; index++ {
		user, createErr := suite.Repo.User.Create(user.ICreate{
			Name:     fmt.Sprintf("Test User %d", index+1),
			Email:    fmt.Sprintf("test%d@example.com", index+1),
			Password: "test123",
			Status:   2,
		})

		if createErr != nil {
			panic(createErr)
		}

		suite.MockUsers[index] = *user
	}
}

func TestUserRepository(test *testing.T) {
	suite.Run(test, new(UserRepoTestSuite))
}
