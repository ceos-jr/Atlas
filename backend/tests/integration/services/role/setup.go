package roleservicetest

import (
	"orb-api/config"
	"orb-api/models"
	"orb-api/repositories/role"
	roleservice "orb-api/services/role"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	Service   *roleservice.Service
	MockRoles []models.Role
}

func (suite *TestSuite) SetupSuite() {
	repo, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Service = roleservice.Setup(&repo.Role)
}

func (suite *TestSuite) TearDownSuite() {
	for index := range suite.MockRoles {
		_, deleteErr := suite.Service.RoleRepo.Delete(role.IDelete{
			RoleID: suite.MockRoles[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}
