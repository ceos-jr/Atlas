package userolerepotest

import (
	"orb-api/config"
	"orb-api/models"
	repository "orb-api/repositories"
	"orb-api/repositories/role"
	"orb-api/repositories/user"
	userrole "orb-api/repositories/user_role"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
)

const (
	MockArraySize = 2
)

type UserRoleRepoTestSuite struct {
	suite.Suite
	Repo         *repository.Repository
	MockUser     []models.User
	MockRole     []models.Role
	MockUserRole []models.UserRole
}

func (suite *UserRoleRepoTestSuite) SetupSuite() {
	repo, setupErr := config.SetupDB("../../.env")

	if setupErr != nil {
		panic(setupErr)
	}

	suite.Repo = repo
	suite.MockUser = make([]models.User, 3)
	suite.MockRole = make([]models.Role, MockArraySize)
	suite.MockUserRole = make([]models.UserRole, MockArraySize)
	suite.SetupMocks()
}

func (suite *UserRoleRepoTestSuite) SetupMocks() {
	for index := 0; index < 3; index++ {
		createdUser, createError := suite.Repo.User.Create(user.ICreate{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Status:   models.UStatusActive,
			Password: faker.Password(),
		})

		if createError != nil {
			panic(createError)
		}

		suite.MockUser[index] = *createdUser
	}

	for index := 0; index < 2; index++ {
		createdRole, createError := suite.Repo.Role.Create(role.ICreate{
			Name:        faker.Name(),
			Description: faker.Sentence(),
		})

		if createError != nil {
			panic(createError)
		}

		suite.MockRole[index] = *createdRole
	}

	userRole, createError := suite.Repo.UserRole.Create(userrole.ICreate{
		RoleID: suite.MockRole[0].ID,
		UserID: suite.MockUser[0].ID,
	})

	if createError != nil {
		panic(createError)
	}

	suite.MockUserRole[0] = *userRole
}

func (suite *UserRoleRepoTestSuite) TearDownSuite() {
	for index := range suite.MockUser {
		_, deleteErr := suite.Repo.User.Delete(user.IDelete{
			ID: suite.MockUser[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
	for index := range suite.MockRole {
		_, deleteErr := suite.Repo.Role.Delete(role.IDelete{
			RoleID: suite.MockRole[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for index := range suite.MockUserRole {
		_, deleteErr := suite.Repo.UserRole.Delete(userrole.IDelete{
			UserRoleID: suite.MockUserRole[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}

	}
}
