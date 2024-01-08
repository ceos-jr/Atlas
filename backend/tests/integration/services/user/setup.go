package userservicetest

import (
	"orb-api/config"
	"orb-api/models"
	"orb-api/services/user"

	userrepo "orb-api/repositories/user"

	"github.com/stretchr/testify/suite"
)

type TestSuit struct {
	suite.Suite
	Service   *user.Service
	MockUsers []models.User
}

// SetupSuite Executed before all tests
func (suite *TestSuit) SetupSuite() {
	repositories, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Service = user.SetupService(&repositories.User)
	suite.MockUsers = make([]models.User, 3)
	suite.SetupMocks()
}

func (suite *TestSuit) SetupMocks() {
	newUser, createErr := suite.Service.UserRepo.Create(userrepo.ICreate{
		Name:     "Gabrigas",
		Email:    "gabrigas@example.com",
		Password: "mostBeautiful",
		Status:   1,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockUsers[0] = *newUser

	newUser2, createErr2 := suite.Service.UserRepo.Create(userrepo.ICreate{
		Name:     "Gabrigas2",
		Email:    "gabrigas2@example.com",
		Password: "mostBeautiful",
		Status:   2,
	})

	if createErr2 != nil {
		panic(createErr)
	}

	suite.MockUsers[1] = *newUser2

	newUser3, createErr3 := suite.Service.UserRepo.Create(userrepo.ICreate{
		Name:     "Gabrigas3",
		Email:    "gabrigas3@example.com",
		Password: "mostBeautiful",
		Status:   1,
	})

	if createErr3 != nil {
		panic(createErr)
	}

	suite.MockUsers[2] = *newUser3

}

// TearDownSuite Executed after all tests
func (suite *TestSuit) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.Service.UserRepo.Delete(userrepo.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}
