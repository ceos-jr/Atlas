package userrepotest

import (
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/user"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	Repo      *repository.Repository
	MockUsers []models.User
}


func (suite *TestSuite) SetupSuite() {

	repo, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
	suite.MockUsers = make([]models.User, 2)
	suite.SetupMocks()
}

func (suite *TestSuite) SetupMocks() {
	newUser, createErr := suite.Repo.User.Create(user.ICreate{
		Name:     "User 01",
		Email:    "user01@example.com",
		Password: "12345678",
		Status:   1,
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockUsers[0] = *newUser
}

func (suite *TestSuite) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.Repo.User.Delete(user.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}
