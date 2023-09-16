package userservicetest

import (
	"orb-api/config"
	"orb-api/models"
	userservice "orb-api/services/user"

	"orb-api/repositories/user"

	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuit struct {
	suite.Suite
	userservice *userservice.UserService
	MockUsers []models.User
}

// Executed before all tests
func (suite *UserServiceTestSuit) SetupSuite() {
	repo, setupError := config.SetupDB("../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.userservice = userservice.SetupUserService(*repo)
	suite.MockUsers = make([]models.User, 2)
	suite.SetupMocks()
}

func (suite *UserServiceTestSuit) SetupMocks() {
	user, createErr := suite.userservice.UserRepo.Create(user.ICreate{
		Name:     "Gabrigas",
		Email:    "gabrigas@example.com",
		Password: "mostBeautiful",
		Status:   1,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockUsers[0] = *user
}

// Executed after all tests
func (suite *UserServiceTestSuit) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.userservice.UserRepo.Delete(user.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}