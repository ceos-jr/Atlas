package user_roleservicetest

import (
	"orb-api/config"
	"orb-api/models"
	user_rolerepo	"orb-api/repositories/user_role"
	userrepo	"orb-api/repositories/user"
	rolerepo	"orb-api/repositories/role"

	user_roleservice "orb-api/services/user_role"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	Service   *user_roleservice.Service
	MockUsers []models.User
	MockRoles []models.Role
	MockUserRoles []models.UserRole
}

// SetupSuite Executed before all tests
func (suite *TestSuite) SetupSuite(){
	repositories, setupError := config.SetupDB("../../.env")

	if setupError != nil {
		panic(setupError)
	}

	suite.Service = user_roleservice.SetupService(&repositories.User, &repositories.Role, &repositories.UserRole)
	suite.MockUserRoles = make([]models.UserRole, 2)
	suite.MockUsers = make([]models.User, 2)
	suite.MockRoles = make([]models.Role, 2)
	suite.SetupMocks()
}

// setting up the mock UserRole
func (suite *TestSuite) SetupMocks() {

	//criar o mock de user
	newUser, createErr := suite.Service.UserRepo.Create(userrepo.ICreate{
		Name:     "User 01",
		Email:    "user01@example.com",
		Password: "mostBeautiful",
		Status:   1,
	})
	if createErr != nil {
		panic(createErr)
	}
	suite.MockUsers[0] = *newUser

	//Criar o mock de role
	newRole, Err := suite.Service.RoleRepo.Create(rolerepo.ICreate{
		Name:     "role 01",
		Description: "description 01",
	})
	if Err != nil {
		panic(Err)
	}
	suite.MockRoles[0] = *newRole

	//Criação do mock UserRole
	newUserRole, errs := suite.Service.UserRoleRepo.Create(user_rolerepo.ICreate{
		RoleID: suite.MockRoles[0].ID,
		UserID: suite.MockUsers[0].ID,
	})
	if errs != nil{
		panic(errs)
	}

	suite.MockUserRoles[0] = *newUserRole
}

// TearDownSuite Executed after all tests
func (suite *TestSuite) TearDownSuite() {
	for index := range suite.MockUsers {
		_, deleteErr := suite.Service.UserRepo.Delete(userrepo.IDelete{
			ID: suite.MockUsers[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for index := range suite.MockRoles {
		_, deleteErr := suite.Service.UserRepo.Delete(userrepo.IDelete{
			ID: suite.MockRoles[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}

	for index := range suite.MockUserRoles {
		_, deleteErr := suite.Service.UserRoleRepo.Delete(user_rolerepo.IDelete{
			UserRoleID: suite.MockUserRoles[index].ID,
		})

		if deleteErr != nil {
			panic(deleteErr)
		}
	}
}