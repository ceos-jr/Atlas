package userrepotest

import (
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/user"

	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	Repo        *repository.Repository
	MockUsers   []models.User
}

func (suite *UserRepoTestSuite) SetupSuite() {
	repo, setupError := config.SetupDB("../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
}

func (suite *UserRepoTestSuite) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.Repo.User.Delete(user.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}



