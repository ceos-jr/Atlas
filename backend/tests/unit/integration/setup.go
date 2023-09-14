package roleservicetest

import (
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/service/role"
	"orb-api/service/role/implementation.go"

	"github.com/stretchr/testify/suite"
)

type RoleServiceTestSuite struct {
	suite.Suite
	Repo *repository.Repository
	MockRoles []models.Role
}

func (suite *RoleServiceTestSuite) SetupSuite() {
	repo, setupError := config.SetupDB("../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
}

func (suite *RoleServiceTestSuite) TearDownSuite() {
	for index := range suite.MockRoles {
		_, deleteErr := suite.Repo.Role.Delete(role.IDelete{
			RoleID: suite.MockRoles[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}

