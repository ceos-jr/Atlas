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
	repo, setupError := config.SetupDB("../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Repo = repo
	suite.MockRoles = make([]models.Role, 2)
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

func (suite *RoleRepoTestSuite) SetupMocks() {
	role, createErr := suite.Repo.Role.Create(role.ICreate{
		Name:        "Role 0",
		Description: "Description 0",
	})

	if createErr != nil {
		panic(createErr)
	}

	suite.MockRoles[0] = *role
}
