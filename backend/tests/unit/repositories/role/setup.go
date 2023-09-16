package rolerepotest

import (
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/role"

	"github.com/stretchr/testify/suite"
)

type RoleRepoTestSuite struct {
	suite.Suite
	Repo      *repository.Repository
	MockRoles []models.Role
}

func (suite *RoleRepoTestSuite) SetupSuite() {
	repo, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
}

func (suite *RoleRepoTestSuite) TearDownSuite() {
	for index := range suite.MockRoles {
		_, deleteErr := suite.Repo.Role.Delete(role.IDelete{
			RoleID: suite.MockRoles[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}
